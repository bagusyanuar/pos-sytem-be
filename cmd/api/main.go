package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bagusyanuar/pos-sytem-be/config"
	"github.com/bagusyanuar/pos-sytem-be/internal/delivery/http/handler"
	"github.com/bagusyanuar/pos-sytem-be/internal/delivery/http/routes"
	"github.com/bagusyanuar/pos-sytem-be/internal/infrastructure/database"
	"github.com/bagusyanuar/pos-sytem-be/internal/infrastructure/fiber"
	"github.com/bagusyanuar/pos-sytem-be/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	appLogger, err := logger.NewLogger(&cfg.Logger, cfg.App.Environment)
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer appLogger.Sync()

	db, err := database.NewPostgresDB(&cfg.Database)

	if err != nil {
		appLogger.Fatal("Failed to connect to database:", err)
	}

	defer func() {
		if err := database.CloseDB(db); err != nil {
			appLogger.Error("Failed to close database:", err)
		}
	}()

	fiberApp := fiber.NewFiberApp(&cfg.Fiber)
	handlers := handler.NewHandlers(cfg)
	routes.SetupRoutes(fiberApp, handlers)

	// Graceful shutdown
	go func() {
		fmt.Printf("🚀 server is running on port %s\n", cfg.App.Port)
		if err := fiberApp.Listen(":" + cfg.App.Port); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	appLogger.Info("Shutting down server...")
	if err := fiberApp.Shutdown(); err != nil {
		appLogger.Error("Server forced to shutdown:", err)
	}

	appLogger.Info("Server exited")

}
