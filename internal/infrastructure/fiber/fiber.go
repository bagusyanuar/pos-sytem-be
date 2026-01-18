package fiber

import (
	"errors"

	"github.com/bagusyanuar/pos-sytem-be/config"
	"github.com/bagusyanuar/pos-sytem-be/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func NewFiberApp(cfg *config.FiberConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Prefork:      cfg.Prefork,
		BodyLimit:    cfg.BodyLimit,
		ErrorHandler: customErrorHandler,
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	return app
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	switch code {
	case fiber.StatusBadRequest:
		return c.Status(code).JSON(response.BadRequest(err.Error()))
	case fiber.StatusUnauthorized:
		return c.Status(code).JSON(response.Unauthorized(err.Error()))
	case fiber.StatusForbidden:
		return c.Status(code).JSON(response.Forbidden(err.Error()))
	case fiber.StatusNotFound:
		return c.Status(code).JSON(response.NotFound("route"))
	default:
		return c.Status(code).JSON(response.InternalServerError(err))
	}
}
