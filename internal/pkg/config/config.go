package config

import (
	"encoding/json"
	"os"

	sqldbcnf "github.com/FurmanovD/postpackage/pkg/config/sqldb"
)

const (
	// Default config file to use
	DefaultConfigFilename = "config.json"

	// default port to use
	DefaultAPIPort = 8080
	// default log level to use
	DefaultLogLevel = "info"

	// default DB maxConnections to use
	DefaultDBMaxConnections = 10
)

// Config stuct for the service
type Config struct {
	APIPort  int    `json:"port" toml:"port" yaml:"port"`
	LogLevel string `json:"loglevel" toml:"loglevel" yaml:"loglevel"`

	SQLConfig sqldbcnf.SQLDBConfig `json:"sqldb" toml:"sqldb" yaml:"sqldb"`
}

// ParseConfig parses the configuration file.
func ParseConfig(args []string) (*Config, error) {
	// TODO(DF): add config validation
	cfg := &Config{
		APIPort:  DefaultAPIPort,
		LogLevel: DefaultLogLevel,
		SQLConfig: sqldbcnf.SQLDBConfig{
			MaxConnections: DefaultDBMaxConnections,
		},
	}

	cfgFile := DefaultConfigFilename

	if len(args) > 0 {
		cfgFile = args[0]
	}

	cfgFileBytes, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(cfgFileBytes, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.APIPort == 0 {
		cfg.APIPort = DefaultAPIPort
	}

	if cfg.LogLevel == "" {
		cfg.LogLevel = DefaultLogLevel
	}

	return cfg, err
}
