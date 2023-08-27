package auth

import (
	"github.com/bagusyanuar/go-yousee/app/auth/controller"
	"github.com/bagusyanuar/go-yousee/app/auth/repositories"
	"github.com/bagusyanuar/go-yousee/app/auth/service"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Builder struct {
	Database *gorm.DB
	Config   *config.Config
	Router   fiber.Router
}

func NewBuilder(db *gorm.DB, cfg *config.Config, router fiber.Router) Builder {
	return Builder{Database: db, Config: cfg, Router: router}
}

func (b *Builder) Build() {
	authRepository := repositories.NewAuth(b.Database)
	authService := service.NewAuth(authRepository, b.Config.JWT)
	authController := controller.NewAuth(authService, b.Router)

	controllers := []any{
		&authController,
	}

	common.RegisterRoutes(controllers...)
}
