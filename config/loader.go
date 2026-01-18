package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// Auto read from environment variables
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading .env file: %w", err)
	}

	setDefaults()

	config := &Config{
		App: AppConfig{
			Name:        viper.GetString("APP_NAME"),
			Version:     viper.GetString("APP_VERSION"),
			Environment: viper.GetString("APP_ENV"),
			Port:        viper.GetString("APP_PORT"),
			Debug:       viper.GetBool("APP_DEBUG"),
		},
		Fiber: FiberConfig{
			ReadTimeout:  time.Duration(viper.GetInt("FIBER_READ_TIMEOUT")) * time.Second,
			WriteTimeout: time.Duration(viper.GetInt("FIBER_WRITE_TIMEOUT")) * time.Second,
			Prefork:      viper.GetBool("FIBER_PREFORK"),
			BodyLimit:    viper.GetInt("FIBER_BODY_LIMIT"),
		},
		Database: DatabaseConfig{
			Host:            viper.GetString("DB_HOST"),
			Port:            viper.GetString("DB_PORT"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			DBName:          viper.GetString("DB_NAME"),
			SSLMode:         viper.GetString("DB_SSLMODE"),
			MaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
			MaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
			ConnMaxLifetime: time.Duration(viper.GetInt("DB_CONN_MAX_LIFETIME")) * time.Second,
		},
	}
	return config, nil
}

func setDefaults() {
	// App defaults
	viper.SetDefault("APP_NAME", "POS System")
	viper.SetDefault("APP_VERSION", "1.0.0")
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("APP_DEBUG", true)

	// Database defaults
	viper.SetDefault("DB_MAX_OPEN_CONNS", 25)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 5)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", 3600)
	viper.SetDefault("DB_SSLMODE", "disable")

	// Redis defaults
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("REDIS_POOL_SIZE", 10)
	viper.SetDefault("REDIS_MIN_IDLE_CONNS", 3)
	viper.SetDefault("REDIS_PASSWORD", "")

	// MinIO defaults
	viper.SetDefault("MINIO_USE_SSL", false)
	viper.SetDefault("MINIO_LOCATION", "us-east-1")

	// JWT defaults
	viper.SetDefault("JWT_EXPIRE_DURATION", 86400)   // 24 hours
	viper.SetDefault("JWT_REFRESH_DURATION", 604800) // 7 days

	// Fiber defaults
	viper.SetDefault("FIBER_READ_TIMEOUT", 10)
	viper.SetDefault("FIBER_WRITE_TIMEOUT", 10)
	viper.SetDefault("FIBER_PREFORK", false)
	viper.SetDefault("FIBER_BODY_LIMIT", 4194304) // 4MB
}
