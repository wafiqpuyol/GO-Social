package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (p *PostStore) CreatePost(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, user_id, title, tags)
		VALUES ($1, $2, $3, $4) RETURNING id,created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeOutDuration)
	defer cancel()

	err := p.db.
		QueryRowContext(ctx, query, post.Content, post.UserID, post.Title, pq.Array(post.Tags)).
		Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	return err
}
