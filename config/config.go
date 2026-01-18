package config

import "time"

type Config struct {
	App      AppConfig
	Fiber    FiberConfig
	Database DatabaseConfig
	Logger   LoggerConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Environment string
	Port        string
	Debug       bool
}

type FiberConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Prefork      bool
	BodyLimit    int
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	LogMode         bool
}

type LoggerConfig struct {
	Level      string
	OutputPath string
	ErrorPath  string
}
