package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Database    DatabaseConfig
	JWT         JWTConfig
	Redis       RedisConfig
	Enviornment string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type JWTConfig struct {
	Secret             string
	AccessTokenLife    time.Duration
	RefreshTokenLife   time.Duration
	RefreshTokenSecret string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func Load() *Config {
	env := os.Getenv("GO_ENV")

	if env == "" {
		env = "development"
	}

	accessTokenLife := 3600
	if val := os.Getenv("ACCESSTOKEN_LIFETIME"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			accessTokenLife = parsed
		}
	}

	refreshTokenLife := 604800
	if val := os.Getenv("REFRESHTOKEN_LIFETIME"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			refreshTokenLife = parsed
		}
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},

		JWT: JWTConfig{
			Secret:             os.Getenv("JWT_SECRET"),
			AccessTokenLife:    time.Duration(accessTokenLife) * time.Second,
			RefreshTokenLife:   time.Duration(refreshTokenLife) * time.Second,
			RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
		},

		Redis: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},

		Enviornment: env,
	}
}
