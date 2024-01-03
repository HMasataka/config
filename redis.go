package config

import "github.com/caarlos0/env/v9"

type RedisConfig struct {
	Host string `json:"host" env:"REDIS_HOST" envDefault:"localhost"`
	Port int    `json:"port" env:"REDIS_PORT" envDefault:"6379"`
}

func NewRedisConfig() (*RedisConfig, error) {
	cfg := &RedisConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
