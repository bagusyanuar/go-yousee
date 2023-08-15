package controller

import (
	"errors"
	"fmt"

	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/app/admin/service"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type City struct {
	cityService service.CityService
	router      fiber.Router
}

func NewCity(citySvc service.CityService, r fiber.Router) City {
	return City{
		cityService: citySvc,
		router:      r,
	}
}

func (c *City) Create(ctx *fiber.Ctx) error {
	var request request.CityRequest

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

	data, err := c.cityService.Create(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error : %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.APIResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (c *City) GetData(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	page := ctx.QueryInt("page", 1)
	perPage := ctx.QueryInt("per_page", 0)
	result, err := c.cityService.GetData(name, page, perPage)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	meta := common.PaginationMeta{
		Limit:      result.Limit,
		Page:       result.Page,
		TotalRows:  result.TotalRows,
		TotalPages: result.TotalPages,
	}
	return ctx.Status(fiber.StatusOK).JSON(common.APIResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result.Rows,
		Meta: map[string]interface{}{
			"pagination": meta,
		},
	})
}

func (c *City) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var request request.CityRequest

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

	data, err := c.cityService.Patch(id, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error : %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.APIResponse{
		Code:    fiber.StatusCreated,
		Message: "success",
		Data:    data,
	})
}

func (c *City) GetDataByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result, err := c.cityService.GetDataByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(common.APIResponse{
				Code:    fiber.StatusNotFound,
				Message: "data not found",
				Data:    nil,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.APIResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result,
	})
}

func (c *City) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.cityService.Delete(id)
	if err != nil {

		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(common.APIResponse{
		Code:    fiber.StatusOK,
		Message: "success",
	})
}

func (c *City) Routes() {
	group := c.router.Group("/city")
	group.Get("/", c.GetData)
	group.Post("/", c.Create)
	group.Get("/:id", c.GetDataByID)
	group.Patch("/:id/patch", c.Patch)
	group.Delete("/:id/delete", c.Delete)
}
