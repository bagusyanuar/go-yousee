package controller

import (
	"strconv"

	"github.com/bagusyanuar/go-yousee/app/admin/service"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/gofiber/fiber/v2"
)

type Province struct {
	provinceService service.ProvinceService
	router          fiber.Router
}

func NewProvince(provinceSvc service.ProvinceService, r fiber.Router) Province {
	return Province{provinceService: provinceSvc, router: r}
}

func (ctrl *Province) Routes() {
	group := ctrl.router.Group("/province")
	group.Get("/", ctrl.GetAllData)
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
