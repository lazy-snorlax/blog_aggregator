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
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: DeleteFeed :exec
Delete from feeds where user_id = $1 and url = $2;

-- name: MarkFeedFetched :one
update feeds
set last_fetched_at = now(), updated_at = now()
where id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
select * from feeds
ORDER BY last_fetched_at ASC NULLS FIRST
limit 1;
