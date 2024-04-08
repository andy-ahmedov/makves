package config

import "github.com/kelseyhightower/envconfig"

type Server struct {
	Port int
}

type Postgres struct {
	Username string
	Password string
	Host     string
	Database string
	Port     int
}

type Config struct {
	PostgresDB Postgres
	Srvr       Server
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("postgres", &cfg.PostgresDB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Srvr); err != nil {
		return nil, err
	}

	return cfg, nil
}
