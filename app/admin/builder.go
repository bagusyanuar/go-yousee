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
	cityRepositoy := repositories.NewCity(b.Database)
	typeRepository := repositories.NewType(b.Database)
	vendorRepository := repositories.NewVendor(b.Database)
	itemRepository := repositories.NewItem(b.Database)

	provinceService := service.NewProvince(provinceRepository)
	cityService := service.NewCity(cityRepositoy)
	typeService := service.NewType(typeRepository)
	vendorService := service.NewVendor(vendorRepository)
	itemService := service.NewItem(itemRepository)

	provinceController := controller.NewProvince(provinceService, b.Router)
	cityController := controller.NewCity(cityService, b.Router)
	typeController := controller.NewType(typeService, b.Router)
	vendorController := controller.NewVendor(vendorService, b.Router)
	itemController := controller.NewItem(itemService, b.Router)

	controllers := []any{
		&provinceController,
		&cityController,
		&typeController,
		&vendorController,
		&itemController,
	}

	common.RegisterRoutes(controllers...)
}
