// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package model

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user"(email,
                   password,
                   nickname,
                   gender,
                   role)
VALUES ($1, $2, $3, $4, $5)
    returning id, created_at, updated_at, deleted_at, email, password, nickname, gender, role
`

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Role     int64  `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Password,
		arg.Nickname,
		arg.Gender,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Nickname,
		&i.Gender,
		&i.Role,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :execrows
update "user"
set deleted_at =$2
where id = $1
  and deleted_at is null
`

type DeleteUserParams struct {
	ID        int64        `json:"id"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteUser, arg.ID, arg.DeletedAt)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, deleted_at, email, password, nickname, gender, role
FROM "user"
WHERE email = $1
    LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Nickname,
		&i.Gender,
		&i.Role,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, updated_at, deleted_at, email, password, nickname, gender, role
FROM "user"
WHERE id = $1
    LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Nickname,
		&i.Gender,
		&i.Role,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
select id, created_at, updated_at, deleted_at, email, password, nickname, gender, role
from "user"
where deleted_at IS NULL
order by id
    limit $1 offset $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Email,
			&i.Password,
			&i.Nickname,
			&i.Gender,
			&i.Role,
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

const updateUser = `-- name: UpdateUser :one
update "user"
set updated_at = $1,
    nickname   = $2,
    gender     = $3,
    role       = $4,
    password   = $5
where id = $6
    returning id, created_at, updated_at, deleted_at, email, password, nickname, gender, role
`

type UpdateUserParams struct {
	UpdatedAt time.Time `json:"updated_at"`
	Nickname  string    `json:"nickname"`
	Gender    string    `json:"gender"`
	Role      int64     `json:"role"`
	Password  string    `json:"password"`
	ID        int64     `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.UpdatedAt,
		arg.Nickname,
		arg.Gender,
		arg.Role,
		arg.Password,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Nickname,
		&i.Gender,
		&i.Role,
	)
	return i, err
}
