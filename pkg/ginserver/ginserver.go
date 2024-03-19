package ginserver

import (
	"context"
	"errors"
	"fmt"
	corelog "log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/FurmanovD/postpackage/pkg/log"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinServer
type GinServer interface {
	Engine() *gin.Engine
	ListenAndServe(port int) error
	ServeEx(srv *http.Server) error
	Stop()
	WaitStopped()
}

type ginServer struct {
	ginEngine *gin.Engine

	log log.Logger

	httpServer *http.Server
	startLock  sync.Mutex

	stopWg  sync.WaitGroup
	stopCh  chan struct{}
	context context.Context
	cancel  func()
}

// New server
func NewGinServer(isDebug bool, logger log.Logger) GinServer {
	mode := gin.ReleaseMode
	if isDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	logEntry := logrus.WithFields(logrus.Fields{"scope": "ginserver"})

	gin.DefaultErrorWriter = logEntry.Writer()

	ctx, cancel := context.WithCancel(context.Background())

	s := &ginServer{
		ginEngine: gin.New(),
		log:       logEntry,
		startLock: sync.Mutex{},
		stopWg:    sync.WaitGroup{},
		stopCh:    make(chan struct{}, 1),
		context:   ctx,
		cancel:    cancel,
	}
	if isDebug {
		s.ginEngine.Use(gin.Logger())
	}

	s.ginEngine.Use(func(c *gin.Context) {
		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		}
		c.Next()
		// After the request, be sure there is no body left open. That leaks file handles.
		if b := c.Request.Body; b != nil {
			b.Close()
		}
	})

	// handle any panic by returning HTTP status 500 error
	s.ginEngine.Use(gin.Recovery())

	return s
}

// Engine returns the underlying Gin engine
func (s *ginServer) Engine() *gin.Engine {
	return s.ginEngine
}

// Serve starts listening with most values are default
func (s *ginServer) ListenAndServe(port int) error {
	return s.ServeEx(
		&http.Server{
			Addr:              fmt.Sprintf(":%d", port),
			Handler:           s.ginEngine,
			MaxHeaderBytes:    1 << 20,
			ReadTimeout:       DefTimeoutRead,       // TLS + headers + body arrived
			ReadHeaderTimeout: DefTimeoutReadHeader, // TLS + headers
			WriteTimeout:      DefTimeoutWrite,      // Time Server has to respond (absolute)
			IdleTimeout:       DefTimeoutIdle,       // Keepalives H1
		})
}

// ServeEx allows to tune an HTTP server configuration
func (s *ginServer) ServeEx(srv *http.Server) error {
	if srv == nil {
		return errors.New("invalid http server")
	}
	s.startLock.Lock()
	defer s.startLock.Unlock()

	if s.httpServer != nil {
		return errors.New("http server starts only once")
	}
	s.httpServer = srv

	s.stopWg.Add(1)
	var err error
	// Serving routine
	go func() {
		defer s.stopWg.Done()

		c := make(chan os.Signal, 3)
		signal.Notify(c, syscall.SIGTERM)
		select {
		case <-c:
		case <-s.stopCh:
		}

		s.httpServer.SetKeepAlivesEnabled(false)
		if err = s.httpServer.Shutdown(context.Background()); err != nil {
			s.log.Errorf("error during server shutdown: %v", err)
		}
		s.cancel()
	}()

	gracehttp.SetLogger(
		corelog.New(
			corelog.Writer(),
			"gracehttp",
			0,
		),
	)

	err = gracehttp.Serve(s.httpServer)
	if err != nil {
		s.log.Errorf("error running http server on address %s: %v", s.httpServer.Addr, err)
		s.Stop()
	}

	s.WaitStopped()

	s.httpServer = nil
	return err
}

// Stop sends stop signal to the server listening routine
func (s *ginServer) Stop() {
	s.stopCh <- struct{}{}
}

// WaitStopped waits to stop the server
func (s *ginServer) WaitStopped() {
	s.stopWg.Wait()
}
