package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerPort int    `env:"SERVER_PORT"`
	Difficult int `env:"CHALLENGE_DIFFICULT"`
	RedisAddress string  `env:"REDIS_ADDRESS"`
	RedisExp int `env:"REDIS_EXP"`
}

func (config *Config) ReadFromEnv() error {
	// Unmarshal config from ENV
	err := env.Parse(config)
	if err != nil {
		return err
	}

	return nil
}

func (config *Config) Address() string {
	return fmt.Sprintf(":%d", config.ServerPort)
}

func (config *Config) RedisExpDur() time.Duration {
	return time.Duration(config.RedisExp) * time.Second;
}