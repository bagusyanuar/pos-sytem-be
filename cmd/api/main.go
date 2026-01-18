package main

import (
	"log"

	"github.com/bagusyanuar/pos-sytem-be/config"
	"github.com/bagusyanuar/pos-sytem-be/internal/infrastructure/fiber"
	"github.com/bagusyanuar/pos-sytem-be/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize logger
	zapLog, err := logger.NewLogger(cfg.App.Environment)
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer zapLog.Sync()

	fiber.NewFiberApp(&cfg.Fiber)
}
