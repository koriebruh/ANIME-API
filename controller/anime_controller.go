package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AnimeController interface {
	GetAnimeById(ctx *fiber.Ctx) error
	GetAllAnime(ctx *fiber.Ctx) error
	AddFavorite(ctx *fiber.Ctx) error
	RemoveFavorite(ctx *fiber.Ctx) error
}
