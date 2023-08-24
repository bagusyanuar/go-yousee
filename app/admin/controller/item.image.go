package controller

import (
	"fmt"

	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/app/admin/service"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/go-playground/validator/v10"
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

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error : %s", err.Error()),
		})
	}

	// validate form request
	validate := validator.New()
	v := &common.CustomValidator{
		Validator: validate,
	}

	if errs := v.Validate(&request); len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.APIResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid request",
			Data:    errs,
		})
	}

	//getting multipart form file
	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["image"]
		for _, file := range files {
			request.Image = file
		}
	}

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
