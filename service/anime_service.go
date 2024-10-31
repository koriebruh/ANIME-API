package service

import (
	"github.com/gofiber/fiber/v2"
	"insert_DM/domain/dto"
)

type AnimeService interface {
	GetAnimeById(ctx *fiber.Ctx) dto.WebRes
	GetAllAnime(ctx *fiber.Ctx) dto.WebRes
	AddFavorite(ctx *fiber.Ctx) dto.WebRes
	RemoveFavorite(ctx *fiber.Ctx) dto.WebRes
}
