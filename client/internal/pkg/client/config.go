package client

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerPort int    `env:"SERVER_PORT"`
	ServerHost string `env:"SERVER_HOST"`
	RequestInterval int `env:"REQUEST_INTERVAL"`
	serverAddress string
}

func (config *Config) ReadFromEnv() error {
	// Unmarshal config from ENV
	err := env.Parse(config)
	if err != nil {
		return err
	}

	return nil
}

func (config *Config) ServerAddress() string {
	if len(config.serverAddress) == 0 {
	  config.serverAddress = fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
	}

	return config.serverAddress
}

func (config *Config) RequestIntervalDur() time.Duration {
	return time.Duration(config.RequestInterval) * time.Second;
}