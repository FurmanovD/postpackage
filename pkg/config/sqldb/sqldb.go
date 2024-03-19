package sqldb

// DBConfig contains parameters to connect to an SQL DB
type SQLDBConfig struct {
	Host           string `json:"host" toml:"host" yaml:"host"`
	Port           int    `json:"port" toml:"port" yaml:"port"`
	User           string `json:"user" toml:"user" yaml:"user"`
	Password       string `json:"password" toml:"password" yaml:"password"`
	Database       string `json:"database" toml:"database" yaml:"database"`
	MaxConnections int    `json:"max_connections" toml:"max_connections" yaml:"max_connections"`
}
