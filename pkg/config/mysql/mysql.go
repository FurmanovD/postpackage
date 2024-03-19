package mysql

import (
	"strconv"

	"github.com/FurmanovD/postpackage/pkg/config/sqldb"
)

const (
	DefaultMySQLHost = "127.0.0.1"
	DefaultMySQLPort = 3306
)

// MySQLConnectionString constructs a MySQL connection string from configuration parameters
func MySQLConnectionString(c sqldb.SQLDBConfig) string {
	// port, 3306 by default
	p := c.Port
	if p <= 0 {
		p = DefaultMySQLPort
	}

	// host, 127.0.0.1 by default
	h := c.Host
	if h == "" {
		h = DefaultMySQLHost
	}
	return c.User + ":" + c.Password + "@tcp(" + h + ":" + strconv.Itoa(p) + ")/" + c.Database
}
