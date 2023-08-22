package controller

import (
	"fmt"

	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/app/admin/service"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/gofiber/fiber/v2"
)

type ItemImage struct {
	itemImageService service.ImageItemService
	router           fiber.Router
}

func NewItemImage(itemImageSvc service.ImageItemService, r fiber.Router) ItemImage {
	return ItemImage{
		itemImageService: itemImageSvc,
		router:           r,
	}
}

func (c *ItemImage) Create(ctx *fiber.Ctx) error {
	var request request.ItemImageRequest
	d, err := c.itemImageService.Create(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error : %s", err.Error()),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(common.APIResponse{
		Code:    fiber.StatusCreated,
		Message: "success",
		Data:    d,
	})
}

func (c *ItemImage) Routes() {
	group := c.router.Group("/item-image")
	group.Post("/", c.Create)
}
