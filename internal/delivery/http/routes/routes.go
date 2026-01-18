package routes

import (
	"github.com/bagusyanuar/pos-sytem-be/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handlers *handler.Handlers) {
	app.Get("/", handlers.Welcome.Welcome)
}
