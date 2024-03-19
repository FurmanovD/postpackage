package webapp

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/FurmanovD/postpackage/internal/app/service"
	"github.com/FurmanovD/postpackage/pkg/ginserver"
	"github.com/FurmanovD/postpackage/pkg/log"
)

// webServer implements the WebServer interface
type webServer struct {
	serviceFacade service.PostService
	ginServer     ginserver.GinServer

	routerGroupAPIV1 *gin.RouterGroup
}

// NewServer creates a WebServer interface instance
func NewServer(serviceFacade service.PostService, logger log.Logger) WebServer {
	srv := ginserver.NewGinServer(true, logger)
	return &webServer{
		serviceFacade:    serviceFacade,
		ginServer:        srv,
		routerGroupAPIV1: srv.Engine().Group("/api").Group("/v1"),
	}
}

// ListenAndServe starts routes serve on the given port
func (s *webServer) ListenAndServe(port int) error {
	if s.ginServer == nil {
		return errors.New("underlying gin server has not been initialized")
	}

	return s.ginServer.ListenAndServe(port)
}
