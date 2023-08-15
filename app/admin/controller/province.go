package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/app/admin/service"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Province struct {
	provinceService service.ProvinceService
	router          fiber.Router
}

func NewProvince(provinceSvc service.ProvinceService, r fiber.Router) Province {
	return Province{provinceService: provinceSvc, router: r}
}

// province endpoint route
func (ctrl *Province) Routes() {
	group := ctrl.router.Group("/province")
	group.Get("/", ctrl.GetAllData)
	group.Post("/", ctrl.Create)
	group.Get("/:id", ctrl.GetDataByID)
	group.Patch("/:id/patch", ctrl.Patch)
	group.Delete("/:id/delete", ctrl.Delete)
}

func (ctrl *Province) GetAllData(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	page, _ := strconv.Atoi(ctx.Query("page"))
	perPage, _ := strconv.Atoi(ctx.Query("per_page"))
	result, err := ctrl.provinceService.GetData(name, page, perPage)
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

func (ctrl *Province) Create(ctx *fiber.Ctx) error {

	var request request.ProvinceRequest

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

	data, err := ctrl.provinceService.Create(request)
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

func (c *Province) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var request request.ProvinceRequest

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

	data, err := c.provinceService.Patch(id, request)
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

func (c *Province) GetDataByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result, err := c.provinceService.GetDataByID(id)
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

func (c *Province) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.provinceService.Delete(id)
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
