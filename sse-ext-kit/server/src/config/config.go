package config

import (
	"github.com/kelseyhightower/envconfig"
	"gitlab.globerce.com/freedom-business/libs/shared-libs/errors"
)

const ChannelDefaultSize = 2

type Config struct {
	PostgresConfig
	GRPCConfig
}

type GRPCConfig struct {
	ListenAddr string `envconfig:"LISTEN_ADDR" default:":9090"`
}

type PostgresConfig struct {
	Host         string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port         string `envconfig:"POSTGRES_PORT" default:"5432"`
	DatabaseName string `envconfig:"POSTGRES_NAME" default:"company"`
	Username     string `envconfig:"POSTGRES_USER" default:"postgres"`
	Password     string `envconfig:"POSTGRES_PASS" default:"mysecretpassword"`
	MaxConns     int    `envconfig:"POSTGRES_MAX_CONNS" default:"30"`
	Schema       string `envconfig:"POSTGRES_SCHEMA"`
}

func InitConfigs() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("company-api", &cfg); err != nil {
		return nil, errors.ENVReadError.SetDevMessage(err.Error())
	}

	return &cfg, nil
}
