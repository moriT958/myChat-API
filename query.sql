-- name: GetThreadIdByUuid :one
SELECT id FROM threads WHERE uuid = $1;

-- name: CreatePost :exec
INSERT INTO posts (uuid, body, thread_id, created_at) VALUES ($1, $2, $3, $4);

-- name: GetPostsByThreadId :many
SELECT uuid, body, created_at FROM posts WHERE thread_id = $1;

