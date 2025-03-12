-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name as feedname, feeds.url, users.name as username
FROM feeds INNER JOIN users 
ON feeds.user_id = users.id;

-- name: GetFeedFromUrl :one
SELECT feeds.id as feed_id, feeds.url as feed_url, feeds.name as feed_name
FROM feeds
WHERE feeds.url = $1;

-- name: GetFeedIDFromUrl :one
SELECT feeds.id as feed_id
FROM feeds
WHERE feeds.url = $1;

-- name: GetFeedID :one
SELECT id FROM feeds
WHERE name = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched = $1, updated_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched ASC NULLS FIRST
LIMIT 1;