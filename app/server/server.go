package server

import (
	"fmt"
	"net/http"

	"github.com/bagusyanuar/go-yousee/app/admin"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func Listen(cfg *config.Config, db *gorm.DB) {
	app := fiber.New(fiber.Config{
		TrustedProxies: []string{"127.0.0.1", "localhost"},
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
			"app_name":    "go-yousee-backend",
			"app_version": "1.0.0",
		})
	})

	api := app.Group("/api")
	adminBuilder := admin.NewBuilder(db, cfg, api)
	adminBuilder.Build()
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(common.APIResponse{
			Code:    http.StatusNotFound,
			Message: "route not found",
			Data:    nil,
		})
	})
	port := fmt.Sprintf(":%s", cfg.Port)
	app.Listen(port)
}
