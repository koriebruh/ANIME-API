package service

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"insert_DM/cnf"
	"insert_DM/domain/dto"
	"insert_DM/repository"
	"insert_DM/utils"
	"net/http"
	"strconv"
)

type AnimeServiceImpl struct {
	*sql.DB
	repository.AnimeRepository
}

func NewAnimeService(DB *sql.DB, animeRepository repository.AnimeRepository) *AnimeServiceImpl {
	return &AnimeServiceImpl{DB: DB, AnimeRepository: animeRepository}
}

func (service AnimeServiceImpl) GetAnimeById(ctx *fiber.Ctx) dto.WebRes {
	var responses dto.WebRes

	// Ambil anime_id dari parameter URL
	animeID, err := strconv.Atoi(ctx.Params("anime_id"))
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusBadRequest,
			Status: "BadRequest",
			Data:   "Invalid anime ID",
		}
		return responses
	}

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer tx.Rollback()

	anime, err := service.AnimeRepository.GetAnimeById(ctx.Context(), tx, animeID)
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusNotFound,
			Status: "NotFound",
			Data:   err.Error(),
		}
		return responses
	}

	err = tx.Commit()
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   err.Error(),
		}
		return responses
	}

	responses = dto.WebRes{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   anime,
	}

	return responses

}

func (service AnimeServiceImpl) GetAllAnime(ctx *fiber.Ctx) dto.WebRes {
	var responses dto.WebRes

	// TAKE TOKEN
	token := ctx.Cookies("token")

	claims := &cnf.JWTClaim{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cnf.JWT_KEY), nil
	})

	// VERIVIKASI JWT
	if err != nil || !tkn.Valid {
		responses = dto.WebRes{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "Invalid token",
		}
		return responses
	}

	// AMBIL USERID YG DI SIMPAN SAAT LOGIN DI JWT
	userID := claims.UserID
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer tx.Rollback()

	favoriteAnimes, err := service.AnimeRepository.GetAllFavorite(ctx.Context(), tx, userID)
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   err.Error(),
		}
		return responses
	}

	err = tx.Commit()
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   err.Error(),
		}
		return responses
	}

	responses = dto.WebRes{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   favoriteAnimes,
	}

	return responses

}

func (service AnimeServiceImpl) AddFavorite(ctx *fiber.Ctx) dto.WebRes {
	var responses dto.WebRes

	userID, _ := strconv.Atoi(ctx.Locals("userID").(string))
	animeID, err := strconv.Atoi(ctx.Params("anime_id"))
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusBadRequest,
			Status: "BadRequest",
			Data:   "Invalid anime ID",
		}
		return responses
	}

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer tx.Rollback()

	err = service.AnimeRepository.AddFavorite(ctx.Context(), tx, userID, animeID)
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   err.Error(),
		}
		return responses
	}

	err = tx.Commit()
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   "Failed to commit transaction",
		}
		return responses
	}

	responses = dto.WebRes{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Anime added to favorites",
	}
	return responses

}

func (service AnimeServiceImpl) RemoveFavorite(ctx *fiber.Ctx) dto.WebRes {
	var responses dto.WebRes

	userID, _ := strconv.Atoi(ctx.Locals("userID").(string))
	animeID, err := strconv.Atoi(ctx.Params("anime_id"))
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusBadRequest,
			Status: "BadRequest",
			Data:   "Invalid anime ID",
		}
		return responses
	}

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer tx.Rollback()

	err = service.AnimeRepository.RemoveFavorite(ctx.Context(), tx, userID, animeID)
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   err.Error(),
		}
		return responses
	}

	err = tx.Commit()
	if err != nil {
		responses = dto.WebRes{
			Code:   http.StatusInternalServerError,
			Status: "InternalServerError",
			Data:   "Failed to commit transaction",
		}
		return responses
	}

	responses = dto.WebRes{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Anime added to favorites",
	}

	return responses

}
