package handler

import (
	"github.com/bagusyanuar/pos-sytem-be/config"
	"github.com/bagusyanuar/pos-sytem-be/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type (
	WelcomHandler interface {
		Welcome(ctx *fiber.Ctx) error
	}

	welcomeHandlerImpl struct {
		Config *config.Config
	}
)

// Welcome implements WelcomHandler.
func (w *welcomeHandlerImpl) Welcome(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(response.Success("success", map[string]any{
		"app_versions": w.Config.App.Version,
		"app_name":     w.Config.App.Name,
	}))
}

func NewWelcomeHandler(config *config.Config) WelcomHandler {
	return &welcomeHandlerImpl{
		Config: config,
	}
}
