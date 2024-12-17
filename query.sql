---------- Queries for User models ----------
-- name: CreateUser :exec
INSERT INTO users (uuid, username, password, created_at) VALUES ($1, $2, $3, $4);

-- name: GetUserByName :one
SELECT uuid, username, password, created_at FROM users WHERE username = $1;

-- name: GetUserByUuid :one
SELECT uuid, username, password, created_at FROM users WHERE uuid = $1;


---------- Queries for Thread models ----------
-- name: CreateThread :exec
INSERT INTO threads (uuid, topic, created_at, user_id) VALUES (
    $1, $2, $3,
    (SELECT users.id FROM users WHERE users.uuid = $4)
);

-- name: GetAllThreads :many
SELECT 
    threads.uuid AS thread_uuid,
    threads.topic,
    threads.created_at AS thread_created_at,
    users.uuid AS user_uuid
FROM
    threads
JOIN
    users ON threads.user_id = users.id
LIMIT $1
OFFSET $2;

-- name: GetThreadByUuid :one
SELECT 
    threads.uuid AS thread_uuid,
    threads.topic,
    threads.created_at AS thread_created_at,
    users.uuid AS user_uuid
FROM
    threads
JOIN
    users ON threads.user_id = users.id
WHERE threads.uuid = $1;

-- name: GetThreadByUserUuid :many
SELECT 
    threads.id,
    threads.uuid AS thread_uuid,
    threads.topic,
    threads.created_at AS thread_created_at,
    users.uuid AS user_uuid
FROM 
    threads 
JOIN 
    users ON threads.user_id = users.id
WHERE
    users.uuid = $1;



---------- Queries for Post models ----------
-- name: CreatePost :exec
INSERT INTO posts (uuid, body, created_at, thread_id, user_id) VALUES (
    $1, $2, $3,
    (SELECT threads.id FROM threads WHERE threads.uuid = $4),
    (SELECT users.id FROM users WHERE users.uuid = $5)
);

-- name: GetPostByUuid :one
SELECT 
    posts.uuid AS post_uuid,
    posts.body,
    posts.created_at AS post_created_at,
    threads.uuid AS thread_uuid,
    users.uuid AS user_uuid
FROM 
    posts 
JOIN threads ON posts.thread_id = threads.id
JOIN users ON posts.user_id = users.id
WHERE
    posts.uuid = $1;


-- name: GetPostByThreadUuid :many
SELECT 
    posts.uuid AS post_uuid,
    posts.body,
    posts.created_at AS post_created_at,
    threads.uuid AS thread_uuid,
    users.uuid AS user_uuid
FROM 
    posts 
JOIN threads ON posts.thread_id = threads.id
JOIN users ON posts.user_id = users.id
WHERE
    threads.uuid = $1;

-- name: GetPostByUserUuid :many
SELECT 
    posts.uuid AS post_uuid,
    posts.body,
    posts.created_at AS post_created_at,
    threads.uuid AS thread_uuid,
    users.uuid AS user_uuid
FROM 
    posts 
JOIN threads ON posts.thread_id = threads.id
JOIN users ON posts.user_id = users.id
WHERE
    users.uuid = $1;
