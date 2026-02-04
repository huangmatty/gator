-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;

-- name: DeleteFeeds :exec
DELETE FROM feeds;