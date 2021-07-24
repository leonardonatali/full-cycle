package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GRPCPort int `envconfig:"grpc_port" default:"50050"`
}

func Load() (cfg *Config, err error) {
	cfg = &Config{}
	err = envconfig.Process("", cfg)
	return
}
