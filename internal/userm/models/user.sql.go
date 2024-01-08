// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  user_id,
  password,
  first_name,
  last_name,
  active,
  email
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING user_id, email, password, first_name, last_name, active, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	UserID    string      `db:"user_id" json:"user_id"`
	Password  string      `db:"password" json:"password"`
	FirstName string      `db:"first_name" json:"first_name"`
	LastName  string      `db:"last_name" json:"last_name"`
	Active    pgtype.Bool `db:"active" json:"active"`
	Email     string      `db:"email" json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserID,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.Active,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUser = `-- name: GetUser :one
SELECT user_id, email, password, first_name, last_name, active, created_at, updated_at, deleted_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
