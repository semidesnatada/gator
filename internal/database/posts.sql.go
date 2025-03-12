// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: posts.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const checkIfPostAlreadyStored = `-- name: CheckIfPostAlreadyStored :one
SELECT EXISTS (SELECT 1
FROM posts
WHERE url = $1
LIMIT 1)
`

func (q *Queries) CheckIfPostAlreadyStored(ctx context.Context, url string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfPostAlreadyStored, url)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getPostsForUser = `-- name: GetPostsForUser :many
SELECT posts.id, posts.title, posts.url, posts.description, posts.published_at FROM posts
INNER JOIN feed_follows
ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetPostsForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

type GetPostsForUserRow struct {
	ID          uuid.UUID
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
}

func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]GetPostsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsForUserRow
	for rows.Next() {
		var i GetPostsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
