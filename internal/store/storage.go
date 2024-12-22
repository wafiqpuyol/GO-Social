package store

import (
	"context"
	"database/sql"
	"time"
)

var (
	queryTimeOutDuration = time.Second * 5
)

type Storage struct {
	Post interface {
		CreatePost(ctx context.Context, post *Post) error
	}
	User interface {
		GetById(ctx context.Context, post *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post: &PostStore{
			db,
		},
		User: &UserStore{
			db,
		},
	}
}
