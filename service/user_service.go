package service

import (
	"github.com/gofiber/fiber/v2"
	"insert_DM/domain/dto"
)

type UserService interface {
	Register(ctx *fiber.Ctx, register dto.UserRegister) dto.WebRes
	Login(ctx *fiber.Ctx, login dto.UserRequestLogin) dto.WebRes
	Logout(ctx *fiber.Ctx) dto.WebRes
}
