package controller

import (
	"github.com/gofiber/fiber/v2"
	"insert_DM/service"
	"net/http"
)

type AnimeControllerImpl struct {
	service.AnimeService
}

func NewAnimeController(animeService service.AnimeService) *AnimeControllerImpl {
	return &AnimeControllerImpl{AnimeService: animeService}
}

func (controller AnimeControllerImpl) GetAnimeById(ctx *fiber.Ctx) error {
	anime := controller.AnimeService.GetAnimeById(ctx)

	return ctx.Status(http.StatusOK).JSON(anime)
}

func (controller AnimeControllerImpl) GetAllAnime(ctx *fiber.Ctx) error {
	animes := controller.AnimeService.GetAllAnime(ctx)

	return ctx.Status(http.StatusOK).JSON(animes)
}

func (controller AnimeControllerImpl) AddFavorite(ctx *fiber.Ctx) error {
	addFav := controller.AnimeService.AddFavorite(ctx)

	return ctx.Status(http.StatusOK).JSON(addFav)
}

func (controller AnimeControllerImpl) RemoveFavorite(ctx *fiber.Ctx) error {
	removeAnime := controller.AnimeService.RemoveFavorite(ctx)

	return ctx.Status(http.StatusOK).JSON(removeAnime)
}
