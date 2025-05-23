package database

import "time"

type CommonDatabaseConfig struct {
	DSN             string        `mapstructure:"dsn"`
	LogLevel        string        `mapstructure:"log_level"`
	SlowThreshold   time.Duration `mapstructure:"slow_threshold"`
	DryRun          bool          `mapstructure:"dry_run"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}
