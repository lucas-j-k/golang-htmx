// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users.queries.sql

package sqlcservice

import (
	"context"
	"database/sql"
	"time"
)

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email, password_hash, row_inserted, row_last_updated FROM user
WHERE id = ?
`

type GetUserRow struct {
	ID             int32
	FirstName      string
	LastName       string
	Email          string
	PasswordHash   string
	RowInserted    time.Time
	RowLastUpdated sql.NullTime
}

func (q *Queries) GetUser(ctx context.Context, id int32) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PasswordHash,
		&i.RowInserted,
		&i.RowLastUpdated,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :many
SELECT id, password_hash, username, first_name, last_name FROM user
WHERE email = ?
`

type GetUserByEmailRow struct {
	ID           int32
	PasswordHash string
	Username     string
	FirstName    string
	LastName     string
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) ([]GetUserByEmailRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserByEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserByEmailRow
	for rows.Next() {
		var i GetUserByEmailRow
		if err := rows.Scan(
			&i.ID,
			&i.PasswordHash,
			&i.Username,
			&i.FirstName,
			&i.LastName,
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

const getUserByUsername = `-- name: GetUserByUsername :many
SELECT id, password_hash, username, first_name, last_name FROM user
WHERE username = ?
`

type GetUserByUsernameRow struct {
	ID           int32
	PasswordHash string
	Username     string
	FirstName    string
	LastName     string
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) ([]GetUserByUsernameRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserByUsernameRow
	for rows.Next() {
		var i GetUserByUsernameRow
		if err := rows.Scan(
			&i.ID,
			&i.PasswordHash,
			&i.Username,
			&i.FirstName,
			&i.LastName,
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

const getUserProfile = `-- name: GetUserProfile :one
SELECT id, username, first_name, last_name, email, row_inserted FROM user
WHERE id = ?
`

type GetUserProfileRow struct {
	ID          int32
	Username    string
	FirstName   string
	LastName    string
	Email       string
	RowInserted time.Time
}

func (q *Queries) GetUserProfile(ctx context.Context, id int32) (GetUserProfileRow, error) {
	row := q.db.QueryRowContext(ctx, getUserProfile, id)
	var i GetUserProfileRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.RowInserted,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :execresult
INSERT INTO user (
  username,
  first_name,
  last_name,
  email,
  password_hash,
  row_inserted,
  row_last_updated
) VALUES (
  ?,
  ?,
  ?,
  ?,
  ?,
  NOW(),
  NULL
)
`

type InsertUserParams struct {
	Username     string
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertUser,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PasswordHash,
	)
}
