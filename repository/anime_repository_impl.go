package repository

import (
	"context"
	"database/sql"
	"fmt"
	"insert_DM/domain"
	"log"
)

type AnimeRepositoryImpl struct {
}

func NewAnimeRepository() *AnimeRepositoryImpl {
	return &AnimeRepositoryImpl{}
}

func (repo AnimeRepositoryImpl) GetAnimeById(ctx context.Context, tx *sql.Tx, animeID int) (domain.AnimeInfo, error) {
	var anime domain.AnimeInfo

	QUERY := `
		SELECT anime_id, name, english_name, other_name, score, genres,
		       synopsis, type, episodes, aired, premiered, status,
		       producers, licensors, studios, source, duration, rating,
		       rank, popularity, favorites, scored_by, members, image_url
		FROM anime_info
		WHERE anime_id = ?
	`

	err := tx.QueryRowContext(ctx, QUERY, animeID).Scan(
		&anime.AnimeID, &anime.Name, &anime.EnglishName, &anime.OtherName, &anime.Score, &anime.Genres,
		&anime.Synopsis, &anime.Type, &anime.Episodes, &anime.Aired, &anime.Premiered, &anime.Status,
		&anime.Producers, &anime.Licensors, &anime.Studios, &anime.Source, &anime.Duration, &anime.Rating,
		&anime.Rank, &anime.Popularity, &anime.Favorites, &anime.ScoredBy, &anime.Members, &anime.ImageURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return anime, fmt.Errorf("anime with id %d not found", animeID)
		}
		return anime, fmt.Errorf("failed to fetch anime: %v", err)
	}
	log.Printf("data: %v and userid %v", anime, animeID)

	return anime, nil
}

func (repo AnimeRepositoryImpl) GetAllFavorite(ctx context.Context, tx *sql.Tx, userID int) ([]domain.AnimeInfo, error) {
	var favoriteAnimes []domain.AnimeInfo

	// Query dengan JOIN untuk mengambil data lengkap anime favorit pengguna
	QUERY := `
		SELECT ai.anime_id, ai.name, ai.english_name, ai.other_name, ai.score, ai.genres,
			   ai.synopsis, ai.type, ai.episodes, ai.aired, ai.premiered, ai.status,
			   ai.producers, ai.licensors, ai.studios, ai.source, ai.duration, ai.rating,
			   ai.rank, ai.popularity, ai.favorites, ai.scored_by, ai.members, ai.image_url
		FROM user_favorites uf
		JOIN anime_info ai ON uf.anime_id = ai.anime_id
		WHERE uf.user_id = ?
	`

	rows, err := tx.QueryContext(ctx, QUERY, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch favorite animes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var anime domain.AnimeInfo
		if err := rows.Scan(
			&anime.AnimeID, &anime.Name, &anime.EnglishName, &anime.OtherName, &anime.Score, &anime.Genres,
			&anime.Synopsis, &anime.Type, &anime.Episodes, &anime.Aired, &anime.Premiered, &anime.Status,
			&anime.Producers, &anime.Licensors, &anime.Studios, &anime.Source, &anime.Duration, &anime.Rating,
			&anime.Rank, &anime.Popularity, &anime.Favorites, &anime.ScoredBy, &anime.Members, &anime.ImageURL,
		); err != nil {
			return nil, fmt.Errorf("failed to scan anime info: %v", err)
		}
		favoriteAnimes = append(favoriteAnimes, anime)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return favoriteAnimes, nil
}

func (repo AnimeRepositoryImpl) AddFavorite(ctx context.Context, tx *sql.Tx, userID int, animeID int) error {
	log.Printf("userid %v and animeid %v", userID, animeID)

	QUERY := "INSERT INTO user_favorites (user_id,anime_id) values (?,?)"
	_, err := tx.ExecContext(ctx, QUERY,
		userID,
		animeID,
	)
	if err != nil {
		return fmt.Errorf("failed to add favorite: %v", err)
	}

	return nil
}

func (repo AnimeRepositoryImpl) RemoveFavorite(ctx context.Context, tx *sql.Tx, userID int, animeID int) error {
	QUERY := "DELETE FROM user_favorites WHERE user_id = ? AND anime_id = ?"
	_, err := tx.QueryContext(ctx, QUERY,
		userID,
		animeID,
	)
	if err != nil {
		return fmt.Errorf("failed to remove favorite: %v", err)
	}

	return nil
}
