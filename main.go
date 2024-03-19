package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"

	"github.com/FurmanovD/postpackage/internal/app/service"
	"github.com/FurmanovD/postpackage/internal/app/webapp"
	"github.com/FurmanovD/postpackage/internal/pkg/config"
	dbpkg "github.com/FurmanovD/postpackage/internal/pkg/db"
	mysqlcnf "github.com/FurmanovD/postpackage/pkg/config/mysql"
	sqldbcnf "github.com/FurmanovD/postpackage/pkg/config/sqldb"
	"github.com/FurmanovD/postpackage/pkg/log"
	"github.com/FurmanovD/postpackage/pkg/sqldb"
)

const (
	// program's exit codes
	errCodeConfigError       = 1
	errCodeDBError           = 2
	errCodeWebSrvListenError = 9

	MigrationsDir           = "./db/migrations"
	GooseMigrationDBDialect = "mysql"
)

// Build information
// The actual information will be stored when 'go build' is called from the Docker file
var (
	Version   = "local-dev"
	BuildTime = time.Now().Format(time.RFC3339)
	GitCommit = ""

	buildInfo = ""

	logger log.Logger
)

func init() {
	buildInfo = fmt.Sprintf(
		"Version: %v BuildTime: %v GitCommit: %v",
		Version,
		BuildTime,
		GitCommit,
	)

	logrus.Info(buildInfo)
}

func main() {
	// Parse flags/config file to populate config
	cfg, err := config.ParseConfig(os.Args[1:])
	if err != nil {
		fmt.Printf("Configuration load error: %+v", err)
		os.Exit(errCodeConfigError)
	}

	initLogging(cfg.LogLevel)
	logger.Infof("Logger initialized with LogLevel: %v", cfg.LogLevel)

	dbStorage, err := initDB(sqldb.NewDB(), cfg.SQLConfig)
	if err != nil {
		fmt.Printf("DB error: %+v", err)
		os.Exit(errCodeDBError)
	}

	// Create a service instance that will do all required operations to DB, storages etc.
	logger.Infof("Creating service instance")
	serviceInstance := service.NewService(dbStorage, logger)
	logger.Infof("Starting a web server on port %d", cfg.APIPort)

	// create a web server instances to serve HTTP endpoints
	webServer := webapp.NewServer(serviceInstance, logger)
	webServer.RegisterRoutes()

	err = webServer.ListenAndServe(cfg.APIPort)
	if err != nil {
		logger.Errorf("Web server start failed: %v", err)
		os.Exit(errCodeWebSrvListenError)
	}
}

// initLogging establishes process logging level
func initLogging(logLevel string) {
	// sets the logging level
	level, err := logrus.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	logger = log.Default()
}

func initDB(db sqldb.SqlDB, dbConfig sqldbcnf.SQLDBConfig) (*dbpkg.Storage, error) {
	logger.Info("Creating a DB connection...")

	err := db.Connect(
		sqldb.DriverMySQL,
		mysqlcnf.MySQLConnectionString(dbConfig),
		dbConfig.MaxConnections,
	)
	if err != nil {
		return nil, fmt.Errorf("Creating DB connection failed: %w", err)
	}
	logger.Info("DB connection is up.")

	// apply migrations
	logger.Info("Applying migrations...")
	if err := goose.SetDialect(GooseMigrationDBDialect); err != nil {
		panic(err)
	}

	err = goose.Up(db.Connection().DB, MigrationsDir)
	if err != nil {
		return nil, fmt.Errorf("Applying migrations failed: %w", err)
	}
	logger.Info("migrations are applied.")

	return dbpkg.NewStorage(db.Connection()), nil
}
