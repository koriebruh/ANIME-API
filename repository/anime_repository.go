package repository

import (
	"context"
	"database/sql"
	"insert_DM/domain"
)

type AnimeRepository interface {
	GetAnimeById(ctx context.Context, tx *sql.Tx, animeID int) (domain.AnimeInfo, error)
	GetAllFavorite(ctx context.Context, tx *sql.Tx, userID int) ([]domain.AnimeInfo, error)
	AddFavorite(ctx context.Context, tx *sql.Tx, userID int, animeID int) error
	RemoveFavorite(ctx context.Context, tx *sql.Tx, userID int, animeID int) error
}
