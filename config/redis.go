package config

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisRing struct {
	Addrs        string
	Password     string
	DB           int
	MaxRetries   int
	PoolSize     int
	PoolTimeout  time.Duration
	IdleTimeout  time.Duration
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Options(cfg *RedisRing) *redis.Options {
	return &redis.Options{
		Addr:     cfg.Addrs,
		Password: cfg.Password,
		DB:       cfg.DB,

		MaxRetries: cfg.MaxRetries,

		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,

		PoolSize:    cfg.PoolSize,
		PoolTimeout: cfg.PoolTimeout,
		IdleTimeout: cfg.IdleTimeout,
	}
}
