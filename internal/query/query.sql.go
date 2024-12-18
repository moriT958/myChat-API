// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package query

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :exec
INSERT INTO posts (uuid, body, created_at, thread_id, user_id) VALUES (
    $1, $2, $3,
    (SELECT threads.id FROM threads WHERE threads.uuid = $4),
    (SELECT users.id FROM users WHERE users.uuid = $5)
)
`

type CreatePostParams struct {
	Uuid      uuid.UUID
	Body      string
	CreatedAt time.Time
	Uuid_2    uuid.UUID
	Uuid_3    uuid.UUID
}

// -------- Queries for Post models ----------
func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) error {
	_, err := q.db.ExecContext(ctx, createPost,
		arg.Uuid,
		arg.Body,
		arg.CreatedAt,
		arg.Uuid_2,
		arg.Uuid_3,
	)
	return err
}

const createThread = `-- name: CreateThread :exec
INSERT INTO threads (uuid, topic, created_at, user_id) VALUES (
    $1, $2, $3,
    (SELECT users.id FROM users WHERE users.uuid = $4)
)
`

type CreateThreadParams struct {
	Uuid      uuid.UUID
	Topic     string
	CreatedAt time.Time
	Uuid_2    uuid.UUID
}

// -------- Queries for Thread models ----------
func (q *Queries) CreateThread(ctx context.Context, arg CreateThreadParams) error {
	_, err := q.db.ExecContext(ctx, createThread,
		arg.Uuid,
		arg.Topic,
		arg.CreatedAt,
		arg.Uuid_2,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (uuid, username, password, created_at) VALUES ($1, $2, $3, $4)
`

type CreateUserParams struct {
	Uuid      uuid.UUID
	Username  string
	Password  string
	CreatedAt time.Time
}

// -------- Queries for User models ----------
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Uuid,
		arg.Username,
		arg.Password,
		arg.CreatedAt,
	)
	return err
}

const getAllThreads = `-- name: GetAllThreads :many
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
OFFSET $2
`

type GetAllThreadsParams struct {
	Limit  int32
	Offset int32
}

type GetAllThreadsRow struct {
	ThreadUuid      uuid.UUID
	Topic           string
	ThreadCreatedAt time.Time
	UserUuid        uuid.UUID
}

func (q *Queries) GetAllThreads(ctx context.Context, arg GetAllThreadsParams) ([]GetAllThreadsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllThreads, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllThreadsRow
	for rows.Next() {
		var i GetAllThreadsRow
		if err := rows.Scan(
			&i.ThreadUuid,
			&i.Topic,
			&i.ThreadCreatedAt,
			&i.UserUuid,
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

const getPostByThreadUuid = `-- name: GetPostByThreadUuid :many
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
    threads.uuid = $1
`

type GetPostByThreadUuidRow struct {
	PostUuid      uuid.UUID
	Body          string
	PostCreatedAt time.Time
	ThreadUuid    uuid.UUID
	UserUuid      uuid.UUID
}

func (q *Queries) GetPostByThreadUuid(ctx context.Context, argUuid uuid.UUID) ([]GetPostByThreadUuidRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostByThreadUuid, argUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostByThreadUuidRow
	for rows.Next() {
		var i GetPostByThreadUuidRow
		if err := rows.Scan(
			&i.PostUuid,
			&i.Body,
			&i.PostCreatedAt,
			&i.ThreadUuid,
			&i.UserUuid,
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

const getPostByUserUuid = `-- name: GetPostByUserUuid :many
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
    users.uuid = $1
`

type GetPostByUserUuidRow struct {
	PostUuid      uuid.UUID
	Body          string
	PostCreatedAt time.Time
	ThreadUuid    uuid.UUID
	UserUuid      uuid.UUID
}

func (q *Queries) GetPostByUserUuid(ctx context.Context, argUuid uuid.UUID) ([]GetPostByUserUuidRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostByUserUuid, argUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostByUserUuidRow
	for rows.Next() {
		var i GetPostByUserUuidRow
		if err := rows.Scan(
			&i.PostUuid,
			&i.Body,
			&i.PostCreatedAt,
			&i.ThreadUuid,
			&i.UserUuid,
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

const getPostByUuid = `-- name: GetPostByUuid :one
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
    posts.uuid = $1
`

type GetPostByUuidRow struct {
	PostUuid      uuid.UUID
	Body          string
	PostCreatedAt time.Time
	ThreadUuid    uuid.UUID
	UserUuid      uuid.UUID
}

func (q *Queries) GetPostByUuid(ctx context.Context, argUuid uuid.UUID) (GetPostByUuidRow, error) {
	row := q.db.QueryRowContext(ctx, getPostByUuid, argUuid)
	var i GetPostByUuidRow
	err := row.Scan(
		&i.PostUuid,
		&i.Body,
		&i.PostCreatedAt,
		&i.ThreadUuid,
		&i.UserUuid,
	)
	return i, err
}

const getThreadByUserUuid = `-- name: GetThreadByUserUuid :many
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
    users.uuid = $1
`

type GetThreadByUserUuidRow struct {
	ID              int32
	ThreadUuid      uuid.UUID
	Topic           string
	ThreadCreatedAt time.Time
	UserUuid        uuid.UUID
}

func (q *Queries) GetThreadByUserUuid(ctx context.Context, argUuid uuid.UUID) ([]GetThreadByUserUuidRow, error) {
	rows, err := q.db.QueryContext(ctx, getThreadByUserUuid, argUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetThreadByUserUuidRow
	for rows.Next() {
		var i GetThreadByUserUuidRow
		if err := rows.Scan(
			&i.ID,
			&i.ThreadUuid,
			&i.Topic,
			&i.ThreadCreatedAt,
			&i.UserUuid,
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

const getThreadByUuid = `-- name: GetThreadByUuid :one
SELECT 
    threads.uuid AS thread_uuid,
    threads.topic,
    threads.created_at AS thread_created_at,
    users.uuid AS user_uuid
FROM
    threads
JOIN
    users ON threads.user_id = users.id
WHERE threads.uuid = $1
`

type GetThreadByUuidRow struct {
	ThreadUuid      uuid.UUID
	Topic           string
	ThreadCreatedAt time.Time
	UserUuid        uuid.UUID
}

func (q *Queries) GetThreadByUuid(ctx context.Context, argUuid uuid.UUID) (GetThreadByUuidRow, error) {
	row := q.db.QueryRowContext(ctx, getThreadByUuid, argUuid)
	var i GetThreadByUuidRow
	err := row.Scan(
		&i.ThreadUuid,
		&i.Topic,
		&i.ThreadCreatedAt,
		&i.UserUuid,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT uuid, username, password, created_at FROM users WHERE username = $1
`

type GetUserByNameRow struct {
	Uuid      uuid.UUID
	Username  string
	Password  string
	CreatedAt time.Time
}

func (q *Queries) GetUserByName(ctx context.Context, username string) (GetUserByNameRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, username)
	var i GetUserByNameRow
	err := row.Scan(
		&i.Uuid,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUuid = `-- name: GetUserByUuid :one
SELECT uuid, username, password, created_at FROM users WHERE uuid = $1
`

type GetUserByUuidRow struct {
	Uuid      uuid.UUID
	Username  string
	Password  string
	CreatedAt time.Time
}

func (q *Queries) GetUserByUuid(ctx context.Context, argUuid uuid.UUID) (GetUserByUuidRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByUuid, argUuid)
	var i GetUserByUuidRow
	err := row.Scan(
		&i.Uuid,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
