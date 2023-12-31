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
	mediaTypeRepository := repositories.NewMediaType(b.Database)
	vendorRepository := repositories.NewVendor(b.Database)
	itemRepository := repositories.NewItem(b.Database)
	itemImageRepository := repositories.NewItemImage(b.Database)
	projectRepository := repositories.NewProject(b.Database)
	projectItemRepository := repositories.NewProjectItem(b.Database)

	provinceService := service.NewProvince(provinceRepository)
	cityService := service.NewCity(cityRepositoy)
	mediaTypeService := service.NewMediaType(mediaTypeRepository)
	vendorService := service.NewVendor(vendorRepository)
	itemService := service.NewItem(itemRepository)
	itemImageService := service.NewItemImage(itemImageRepository)
	projectService := service.NewProject(projectRepository)
	projectItemService := service.NewProjectItem(projectItemRepository)

	provinceController := controller.NewProvince(provinceService, b.Router)
	cityController := controller.NewCity(cityService, b.Router)
	mediaTypeController := controller.NewMediaType(mediaTypeService, b.Router)
	vendorController := controller.NewVendor(vendorService, b.Router)
	itemController := controller.NewItem(itemService, b.Router)
	itemImageController := controller.NewItemImage(itemImageService, b.Router)
	projectController := controller.NewProject(projectService, b.Router)
	projectItemController := controller.NewProjectItem(projectItemService, b.Router)

	controllers := []any{
		&provinceController,
		&cityController,
		&mediaTypeController,
		&vendorController,
		&itemController,
		&projectController,
		&projectItemController,
		&itemImageController,
	}

	common.RegisterRoutes(controllers...)
}
