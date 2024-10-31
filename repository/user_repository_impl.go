package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"insert_DM/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) error {
	// Validate If Acc Already ?
	var count int
	err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("username %s already exists", user.Username)
	}

	SQL := "INSERT INTO users (username, password) values (?,?)"
	_, err = tx.ExecContext(ctx, SQL,
		user.Username,
		user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (int, error) {
	var userLogin domain.User

	QUERY := "SELECT user_id, username, password FROM users WHERE username = ?"
	err := tx.QueryRowContext(ctx, QUERY, user.Username).Scan(&userLogin.UserId, &userLogin.Username, &userLogin.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("login failed: incorrect user and password")
		}
		return 0, fmt.Errorf("login failed: %v", err)
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(user.Password))
	if errPass != nil {
		return 0, errors.New("login failed: incorrect password")
	}

	//#MENGEMBALIKAN ID YG AKAN DI SIMPAN DI JWT
	return userLogin.UserId, nil

}
