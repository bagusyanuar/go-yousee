package server

import (
	"fmt"
	"net/http"

	"github.com/bagusyanuar/go-yousee/app/admin"
	"github.com/bagusyanuar/go-yousee/app/auth"
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

	//init endpoint group
	api := app.Group("/api")

	//build app auth scheme
	authBuilder := auth.NewBuilder(db, cfg, api)
	authBuilder.Build()

	//bauild app admin scheme
	adminBuilder := admin.NewBuilder(db, cfg, api)
	adminBuilder.Build()

	//throw route not found
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
