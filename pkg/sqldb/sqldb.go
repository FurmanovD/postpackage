package sqldb

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	// sqlx debug logger
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

const (
	DriverMySQL = "mysql"

	connStringParams = "parseTime=true"
)

// SqlDB interface for db
type SqlDB interface {
	Connect(driver, connStr string, maxConn int) error
	Connection() *sqlx.DB
	ConnectionString() string
}

// sqldb implements the SqlDB interface
type sqldb struct {
	connStr string
	conn    *sqlx.DB
}

func NewDB() SqlDB {
	return &sqldb{}
}

func (d *sqldb) Connect(driver, connStr string, maxConn int) error {
	err := d.createConnection(driver, connStr, connStringParams, maxConn)
	if err != nil {
		return err
	}
	return nil
}

func (d *sqldb) Connection() *sqlx.DB {
	return d.conn
}

func (d *sqldb) ConnectionString() string {
	return d.connStr
}

func (d *sqldb) createConnection(driverName, connStr, params string, maxConn int) error {
	dsn := connStr
	if strings.Contains(dsn, "?") {
		if dsn[len(dsn)-1] != '?' {
			dsn += "&"
		}
		dsn += params
	} else {
		dsn += "?" + params
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return err
	}
	// initiate zerolog
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zlogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	// prepare logger
	loggerOptions := []sqldblogger.Option{
		sqldblogger.WithSQLQueryFieldname("sql"),
		sqldblogger.WithWrapResult(false),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	}
	// wrap *sql.DB to transparent logger
	db = sqldblogger.OpenDriver(dsn, db.Driver(), zerologadapter.New(zlogger), loggerOptions...)
	// pass it sqlx
	conn := sqlx.NewDb(db, driverName)

	conn.SetMaxOpenConns(maxConn)
	conn.SetMaxIdleConns(0)

	d.conn = conn
	d.connStr = connStr

	return nil
}
