package repository

import (
	"context"
	"database/sql"
	"insert_DM/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) error
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (int, error)
}
