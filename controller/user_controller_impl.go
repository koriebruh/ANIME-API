package controller

import (
	"github.com/gofiber/fiber/v2"
	"insert_DM/domain/dto"
	"insert_DM/service"
	"net/http"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{userService: userService}
}

func (controller UserControllerImpl) Register(ctx *fiber.Ctx) error {
	//# Decode : json to obj
	var registerUser dto.UserRegister
	if err := ctx.BodyParser(&registerUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebRes{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	registerResponse := controller.userService.Register(ctx, registerUser)

	return ctx.Status(http.StatusCreated).JSON(registerResponse)
}

func (controller UserControllerImpl) Login(ctx *fiber.Ctx) error {
	var loginUser dto.UserRequestLogin
	if err := ctx.BodyParser(&loginUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebRes{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	loginResponse := controller.userService.Login(ctx, loginUser)

	return ctx.Status(http.StatusOK).JSON(loginResponse)
}

func (controller UserControllerImpl) Logout(ctx *fiber.Ctx) error {

	logout := controller.userService.Logout(ctx)

	return ctx.Status(http.StatusOK).JSON(logout)
}
