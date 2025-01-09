package sqlstore

import "go.signoz.io/signoz/pkg/config"

// Config satisfies the confmap.Config interface
var _ config.Config = (*Config)(nil)

type Config struct {
	Provider string         `mapstructure:"provider"`
	Sqlite   SqliteConfig   `mapstructure:"sqlite"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type SqliteConfig struct {
	Path         string `mapstructure:"path"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func NewConfigFactory() config.ConfigFactory {
	return config.NewConfigFactory(newConfig)
}

func newConfig() config.Config {
	return &Config{
		Provider: "sqlite",
		Sqlite: SqliteConfig{
			Path:         "/var/lib/signoz/signoz.db",
			MaxOpenConns: 10,
		},
	}

}

func (c *Config) Key() string {
	return "sqlstore"
}

func (c *Config) Validate() error {
	return nil
}
