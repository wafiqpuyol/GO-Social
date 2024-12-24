package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
}
type password struct {
	text *string
	hash []byte
}

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) GetById(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password, is_active)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, is_active
	`

	err := u.db.
		QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.IsActive).
		Scan(&user.ID, &user.CreatedAt, &user.IsActive)
	return err
}
