-- name: CreateFeedFollow :one
WITH inserted_row AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
) SELECT
    inserted_row.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_row
INNER JOIN users
ON inserted_row.user_id = users.id
INNER JOIN feeds
ON inserted_row.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name as feed_name, users.name as user_name
FROM feed_follows
INNER JOIN feeds
ON feed_follows.feed_id = feeds.id
INNER JOIN users
ON feed_follows.user_id = users.id
WHERE users.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 
AND feed_id = $2;