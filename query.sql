-- name: CreateUser :exec
INSERT INTO users (uuid, username, password, created_at) VALUES ($1, $2, $3, $4);

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByName :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByUuid :one
SELECT * FROM users WHERE uuid = $1;


-- name: CreateThread :exec
INSERT INTO threads (uuid, topic, created_at, user_id) VALUES ($1, $2, $3, $4);

-- name: GetAllThreads :many
SELECT * FROM threads LIMIT $1 OFFSET $2;

-- name: GetThreadById :one
SELECT * FROM threads WHERE id = $1;

-- name: GetThreadByUuid :one
SELECT * FROM threads WHERE uuid = $1;

-- name: GetThreadByUserId :many
SELECT * FROM threads WHERE user_id = $1;


-- name: CreatePost :exec
INSERT INTO posts (uuid, body, thread_id, created_at, user_id) VALUES ($1, $2, $3, $4, $5);

-- name: GetPostById :one
SELECT * FROM posts WHERE id = $1;

-- name: GetPostByUuid :one
SELECT * FROM posts WHERE uuid = $1;

-- name: GetPostByThreadId :many
SELECT * FROM posts WHERE thread_id = $1;

-- name: GetPostByUserId :many
SELECT * FROM posts WHERE user_id = $1;
