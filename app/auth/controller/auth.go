package controller

import (
	"fmt"

	"github.com/bagusyanuar/go-yousee/app/auth/request"
	"github.com/bagusyanuar/go-yousee/app/auth/service"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	authService service.AuthService
	router      fiber.Router
}

func NewAuth(authSvc service.AuthService, r fiber.Router) Auth {
	return Auth{authService: authSvc, router: r}
}

func (c *Auth) SignIn(ctx *fiber.Ctx) error {
	var request request.AuthRequest

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

	token, err := c.authService.SignIn(request)

	if err != nil {
		switch err {
		case common.ErrPasswordNotMatch:
			return ctx.Status(fiber.StatusUnauthorized).JSON(common.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})
		case common.ErrUserNotFound:
			return ctx.Status(fiber.StatusNotFound).JSON(common.APIResponse{
				Code:    fiber.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Message: fmt.Sprintf("internal server error : %s", err.Error()),
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(common.APIResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    token,
	})
}

func (c *Auth) Routes() {
	group := c.router.Group("/auth")
	group.Post("/sign-in", c.SignIn)
}
