package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestConfig_Address(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
		want   string
	}{
		{"positive", &Config{123, 0, "", 0}, ":123"},
		{"empty config", &Config{}, ":0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.config.Address(); got != tt.want {
				t.Errorf("Config.Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_RedisExpDur(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
		want   time.Duration
	}{
		{"positive", &Config{0, 0, "", 10}, 10 * time.Second},
		{"empty config", &Config{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.config.RedisExpDur(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.RedisExpDur() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ReadFromEnv(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		wantConfig  *Config
		wantErr bool
	}{
		{"empty env", map[string]string{"SERVER_PORT": "", "CHALLENGE_DIFFICULT": "", "REDIS_ADDRESS": "", "REDIS_EXP": "" }, &Config{0, 0, "", 0}, false},
	  {"positive", map[string]string{"SERVER_PORT": "3000", "CHALLENGE_DIFFICULT": "1", "REDIS_ADDRESS": "redis", "REDIS_EXP": "123" }, &Config{3000, 1, "redis", 123}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				if v == "" {
					os.Unsetenv(k)
				} else {
					os.Setenv(k, v)
				}
			}
			conf := Config{}
			if err := conf.ReadFromEnv(); (err != nil) != tt.wantErr {
				t.Errorf("Config.ReadFromEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
			if conf != *tt.wantConfig {
				t.Errorf("Config.ReadFromEnv() conf = %v, wantConf %v", conf, tt.wantConfig)
			}
		})
	}
}
