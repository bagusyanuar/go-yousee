package admin

import (
	"github.com/bagusyanuar/go-yousee/app/admin/controller"
	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/service"
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
	provinceRepository := repositories.NewProvince(b.Database)

	provinceService := service.NewProvince(provinceRepository)

	provinceController := controller.NewProvince(provinceService, b.Router)

	controllers := []any{
		&provinceController,
	}

	common.RegisterRoutes(controllers...)
}
