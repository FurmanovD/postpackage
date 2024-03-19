package mysql

import (
	"strconv"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"

	"github.com/FurmanovD/go-kit/randomstring"
	"github.com/FurmanovD/postpackage/pkg/config/sqldb"
)

func TestMySqlConnectionString(t *testing.T) {

	testCases := map[string]interface{}{
		"DefaultHostOk": caseDefaultHostOk(t),
		"DefaultPortOk": caseDefaultPortOk(t),
		"AllIsFilled":   caseAllIsFilled(t),
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// NOTE(DF): An example of "quick" testing framework usage.
			if err := quick.Check(tc, nil); err != nil {
				t.Errorf("%v case failed with an error: %+v", name, err)
			}
		})
	}
}

// all the test cases to check by testing/quick
func caseDefaultHostOk(t *testing.T) interface{} {
	return func(c sqldb.SQLDBConfig) bool {
		// check default host
		c.Host = ""

		// make sure Port is > 0:
		if c.Port <= 0 {
			c.Port *= -1
		} else if c.Port == 0 {
			c.Port = 123
		}

		testStr := MySQLConnectionString(c)
		return assert.Equal(t,
			c.User+":"+c.Password+"@tcp(127.0.0.1:"+strconv.Itoa(c.Port)+")/"+c.Database,
			testStr,
			"Invalid default host connection string built",
		)
	}
}

func caseDefaultPortOk(t *testing.T) interface{} {
	return func(c sqldb.SQLDBConfig) bool {
		// check default port
		if c.Port > 0 {
			c.Port *= -1
		}

		// make sure Host is not empty
		if c.Host == "" {
			c.Host = randomstring.NonEmptyUTF8Printable(50, nil)
		}

		testStr := MySQLConnectionString(c)
		return assert.Equal(t,
			c.User+":"+c.Password+"@tcp("+c.Host+":"+strconv.Itoa(DefaultMySQLPort)+")/"+c.Database,
			testStr,
			"Invalid default port connection string built")
	}
}

func caseAllIsFilled(t *testing.T) interface{} {
	return func(c sqldb.SQLDBConfig) bool {
		defaultMustBeSet := false

		host := c.Host
		if c.Host == "" {
			defaultMustBeSet = true
			host = DefaultMySQLHost
		}

		port := c.Port
		if c.Port <= 0 {
			defaultMustBeSet = true
			port = DefaultMySQLPort
		}

		testStr := MySQLConnectionString(c)
		expected := c.User + ":" + c.Password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + c.Database

		if defaultMustBeSet {
			return assert.Equal(t,
				expected,
				testStr,
				"Invalid connection string built(default host or port expected)")
		} else {
			return assert.Equal(t,
				expected,
				testStr,
				"Invalid connection string built")
		}
	}
}
