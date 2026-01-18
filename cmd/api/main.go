package main

import (
	"log"

	"github.com/bagusyanuar/pos-sytem-be/config"
	"github.com/bagusyanuar/pos-sytem-be/internal/infrastructure/fiber"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	fiber.NewFiberApp(&cfg.Fiber)
}
