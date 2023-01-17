// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const user_CountAll = `-- name: User_CountAll :one
SELECT      COUNT(*)
FROM        users
`

func (q *Queries) User_CountAll(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, user_CountAll)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const user_Create = `-- name: User_Create :one
INSERT INTO users ( username, email, password )
VALUES      ( $1, $2, $3 )
RETURNING   user_id, created_at, updated_at, is_admin, username, email, password
`

type User_CreateParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) User_Create(ctx context.Context, arg User_CreateParams) (User, error) {
	row := q.db.QueryRowContext(ctx, user_Create, arg.Username, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const user_DeleteWhereIdEquals = `-- name: User_DeleteWhereIdEquals :one
DELETE FROM users
WHERE       user_id = $1
RETURNING   user_id, created_at, updated_at, is_admin, username, email, password
`

func (q *Queries) User_DeleteWhereIdEquals(ctx context.Context, userID uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, user_DeleteWhereIdEquals, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const user_GetAll = `-- name: User_GetAll :many
SELECT      user_id, created_at, username
FROM        users
ORDER BY    created_at DESC
`

type User_GetAllRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetAll(ctx context.Context) ([]User_GetAllRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetAllRow
	for rows.Next() {
		var i User_GetAllRow
		if err := rows.Scan(&i.UserID, &i.CreatedAt, &i.Username); err != nil {
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

const user_GetAllBetweenDates = `-- name: User_GetAllBetweenDates :many
SELECT      user_id, created_at, username
FROM        users
WHERE       created_at BETWEEN SYMMETRIC $1 AND $2
ORDER BY    created_at DESC
`

type User_GetAllBetweenDatesParams struct {
	CreatedAt   time.Time `json:"created_at"`
	CreatedAt_2 time.Time `json:"created_at_2"`
}

type User_GetAllBetweenDatesRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetAllBetweenDates(ctx context.Context, arg User_GetAllBetweenDatesParams) ([]User_GetAllBetweenDatesRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetAllBetweenDates, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetAllBetweenDatesRow
	for rows.Next() {
		var i User_GetAllBetweenDatesRow
		if err := rows.Scan(&i.UserID, &i.CreatedAt, &i.Username); err != nil {
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

const user_GetAllWhereCreatedAt = `-- name: User_GetAllWhereCreatedAt :many
SELECT      user_id, created_at, username
FROM        users
WHERE       created_at = $1
ORDER BY    created_at DESC
`

type User_GetAllWhereCreatedAtRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetAllWhereCreatedAt(ctx context.Context, createdAt time.Time) ([]User_GetAllWhereCreatedAtRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetAllWhereCreatedAt, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetAllWhereCreatedAtRow
	for rows.Next() {
		var i User_GetAllWhereCreatedAtRow
		if err := rows.Scan(&i.UserID, &i.CreatedAt, &i.Username); err != nil {
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

const user_GetAllWithPaginationFirstPage = `-- name: User_GetAllWithPaginationFirstPage :many
SELECT      user_id, created_at, username
FROM        users
ORDER BY    created_at DESC
LIMIT       $1
`

type User_GetAllWithPaginationFirstPageRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetAllWithPaginationFirstPage(ctx context.Context, limit int32) ([]User_GetAllWithPaginationFirstPageRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetAllWithPaginationFirstPage, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetAllWithPaginationFirstPageRow
	for rows.Next() {
		var i User_GetAllWithPaginationFirstPageRow
		if err := rows.Scan(&i.UserID, &i.CreatedAt, &i.Username); err != nil {
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

const user_GetAllWithPaginationNextPage = `-- name: User_GetAllWithPaginationNextPage :many
SELECT      user_id, created_at, username
FROM        users
WHERE
    (
      created_at <= $1::TIMESTAMP
      OR 
      ( created_at = $1::TIMESTAMP AND user_id < $2::uuid )
    )
ORDER BY    created_at DESC
LIMIT       $3::int
`

type User_GetAllWithPaginationNextPageParams struct {
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	Limit     int32     `json:"_limit"`
}

type User_GetAllWithPaginationNextPageRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetAllWithPaginationNextPage(ctx context.Context, arg User_GetAllWithPaginationNextPageParams) ([]User_GetAllWithPaginationNextPageRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetAllWithPaginationNextPage, arg.CreatedAt, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetAllWithPaginationNextPageRow
	for rows.Next() {
		var i User_GetAllWithPaginationNextPageRow
		if err := rows.Scan(&i.UserID, &i.CreatedAt, &i.Username); err != nil {
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

const user_GetSentTransactionsWhereUserIdEquals = `-- name: User_GetSentTransactionsWhereUserIdEquals :many
SELECT      
    users.user_id, users.created_at, users.username,
    transactions.transaction_id, transactions.amount
FROM
    users
    INNER JOIN  transactions
    ON $1::uuid = transactions.sender_id
WHERE 
    ( transactions.created_at <= $2::TIMESTAMP OR 
    ( transactions.created_at =  $2::TIMESTAMP AND users.user_id < $1::uuid ))
ORDER BY    transactions.created_at DESC
LIMIT       $3::int
`

type User_GetSentTransactionsWhereUserIdEqualsParams struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Limit     int32     `json:"_limit"`
}

type User_GetSentTransactionsWhereUserIdEqualsRow struct {
	UserID        uuid.UUID `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	Username      string    `json:"username"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Amount        int32     `json:"amount"`
}

func (q *Queries) User_GetSentTransactionsWhereUserIdEquals(ctx context.Context, arg User_GetSentTransactionsWhereUserIdEqualsParams) ([]User_GetSentTransactionsWhereUserIdEqualsRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetSentTransactionsWhereUserIdEquals, arg.UserID, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetSentTransactionsWhereUserIdEqualsRow
	for rows.Next() {
		var i User_GetSentTransactionsWhereUserIdEqualsRow
		if err := rows.Scan(
			&i.UserID,
			&i.CreatedAt,
			&i.Username,
			&i.TransactionID,
			&i.Amount,
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

const user_GetSentTransactionsWhereUserIdEqualsFirstPage = `-- name: User_GetSentTransactionsWhereUserIdEqualsFirstPage :many
SELECT      
    users.user_id, users.created_at, users.username,
    transactions.transaction_id, transactions.amount
FROM
    users
    INNER JOIN  transactions
    ON $1::uuid = transactions.sender_id
ORDER BY    transactions.created_at DESC
LIMIT       $2::int
`

type User_GetSentTransactionsWhereUserIdEqualsFirstPageParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"_limit"`
}

type User_GetSentTransactionsWhereUserIdEqualsFirstPageRow struct {
	UserID        uuid.UUID `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	Username      string    `json:"username"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Amount        int32     `json:"amount"`
}

func (q *Queries) User_GetSentTransactionsWhereUserIdEqualsFirstPage(ctx context.Context, arg User_GetSentTransactionsWhereUserIdEqualsFirstPageParams) ([]User_GetSentTransactionsWhereUserIdEqualsFirstPageRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetSentTransactionsWhereUserIdEqualsFirstPage, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetSentTransactionsWhereUserIdEqualsFirstPageRow
	for rows.Next() {
		var i User_GetSentTransactionsWhereUserIdEqualsFirstPageRow
		if err := rows.Scan(
			&i.UserID,
			&i.CreatedAt,
			&i.Username,
			&i.TransactionID,
			&i.Amount,
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

const user_GetWhereIdEquals = `-- name: User_GetWhereIdEquals :one
SELECT      user_id, created_at, username
FROM        users
WHERE       user_id = $1
LIMIT 1
`

type User_GetWhereIdEqualsRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetWhereIdEquals(ctx context.Context, userID uuid.UUID) (User_GetWhereIdEqualsRow, error) {
	row := q.db.QueryRowContext(ctx, user_GetWhereIdEquals, userID)
	var i User_GetWhereIdEqualsRow
	err := row.Scan(&i.UserID, &i.CreatedAt, &i.Username)
	return i, err
}

const user_GetWhereUsernameEquals = `-- name: User_GetWhereUsernameEquals :one
SELECT      user_id, created_at, username
FROM        users
WHERE       username = $1
LIMIT 1
`

type User_GetWhereUsernameEqualsRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetWhereUsernameEquals(ctx context.Context, username string) (User_GetWhereUsernameEqualsRow, error) {
	row := q.db.QueryRowContext(ctx, user_GetWhereUsernameEquals, username)
	var i User_GetWhereUsernameEqualsRow
	err := row.Scan(&i.UserID, &i.CreatedAt, &i.Username)
	return i, err
}

const user_GetWhereUsernameIncludes = `-- name: User_GetWhereUsernameIncludes :many
SELECT      user_id, created_at, username
FROM        users
WHERE       username ILIKE '%' || ( $1 ) || '%'
ORDER BY    created_at DESC
`

type User_GetWhereUsernameIncludesRow struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

func (q *Queries) User_GetWhereUsernameIncludes(ctx context.Context, dollar_1 sql.NullString) ([]User_GetWhereUsernameIncludesRow, error) {
	rows, err := q.db.QueryContext(ctx, user_GetWhereUsernameIncludes, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User_GetWhereUsernameIncludesRow
	for rows.Next() {
		var i User_GetWhereUsernameIncludesRow
		if err := rows.Scan(&i.UserID, &i.CreatedAt, &i.Username); err != nil {
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

const user_LoginWithUsername = `-- name: User_LoginWithUsername :one
SELECT      user_id, created_at, updated_at, is_admin, username, email, password
FROM        users
WHERE       username = $1
LIMIT       1
`

func (q *Queries) User_LoginWithUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, user_LoginWithUsername, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const user_UpdateUserToAdminWhereIdEquals = `-- name: User_UpdateUserToAdminWhereIdEquals :one
UPDATE      users
SET         is_admin = TRUE
WHERE       user_id = $1
RETURNING   user_id, created_at, updated_at, is_admin, username, email, password
`

func (q *Queries) User_UpdateUserToAdminWhereIdEquals(ctx context.Context, userID uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, user_UpdateUserToAdminWhereIdEquals, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const user_UpdateUsernameWhereIdEquals = `-- name: User_UpdateUsernameWhereIdEquals :one
UPDATE      users
SET         username = $1
WHERE       user_id = $2
RETURNING   user_id, created_at, updated_at, is_admin, username, email, password
`

type User_UpdateUsernameWhereIdEqualsParams struct {
	Username string    `json:"username"`
	UserID   uuid.UUID `json:"user_id"`
}

func (q *Queries) User_UpdateUsernameWhereIdEquals(ctx context.Context, arg User_UpdateUsernameWhereIdEqualsParams) (User, error) {
	row := q.db.QueryRowContext(ctx, user_UpdateUsernameWhereIdEquals, arg.Username, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const user_UpdateUsernameWhereUsernameEquals = `-- name: User_UpdateUsernameWhereUsernameEquals :one
UPDATE      users
SET         username = $1
WHERE       username = $2
RETURNING   user_id, created_at, updated_at, is_admin, username, email, password
`

type User_UpdateUsernameWhereUsernameEqualsParams struct {
	Username   string `json:"username"`
	Username_2 string `json:"username_2"`
}

func (q *Queries) User_UpdateUsernameWhereUsernameEquals(ctx context.Context, arg User_UpdateUsernameWhereUsernameEqualsParams) (User, error) {
	row := q.db.QueryRowContext(ctx, user_UpdateUsernameWhereUsernameEquals, arg.Username, arg.Username_2)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}
