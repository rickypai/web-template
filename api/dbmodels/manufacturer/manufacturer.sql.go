// Code generated by sqlc. DO NOT EDIT.
// source: manufacturer.sql

package manufacturer

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const countTotal = `-- name: CountTotal :one
SELECT COUNT(id) FROM manufacturers
`

func (q *Queries) CountTotal(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTotal)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOne = `-- name: CreateOne :one
INSERT INTO manufacturers(name, created_at, modified_at) VALUES($1, $2, $3) RETURNING id, name, created_at, modified_at
`

type CreateOneParams struct {
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (q *Queries) CreateOne(ctx context.Context, arg CreateOneParams) (Manufacturer, error) {
	row := q.db.QueryRowContext(ctx, createOne, arg.Name, arg.CreatedAt, arg.ModifiedAt)
	var i Manufacturer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getByID = `-- name: GetByID :one
SELECT id, name, created_at, modified_at FROM manufacturers WHERE id = $1 LIMIT 1
`

func (q *Queries) GetByID(ctx context.Context, id int64) (Manufacturer, error) {
	row := q.db.QueryRowContext(ctx, getByID, id)
	var i Manufacturer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getManyByIDs = `-- name: GetManyByIDs :many
SELECT id, name, created_at, modified_at FROM manufacturers WHERE id = ANY($1::bigint[])
`

func (q *Queries) GetManyByIDs(ctx context.Context, dollar_1 []int64) ([]Manufacturer, error) {
	rows, err := q.db.QueryContext(ctx, getManyByIDs, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Manufacturer
	for rows.Next() {
		var i Manufacturer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.ModifiedAt,
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

const listByPattern = `-- name: ListByPattern :many
SELECT id, name, created_at, modified_at FROM manufacturers WHERE name LIKE $1 ORDER BY name ASC LIMIT $2
`

type ListByPatternParams struct {
	Name  string
	Limit int32
}

func (q *Queries) ListByPattern(ctx context.Context, arg ListByPatternParams) ([]Manufacturer, error) {
	rows, err := q.db.QueryContext(ctx, listByPattern, arg.Name, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Manufacturer
	for rows.Next() {
		var i Manufacturer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.ModifiedAt,
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

const listOffset = `-- name: ListOffset :many
SELECT id, name, created_at, modified_at FROM manufacturers LIMIT $1 OFFSET $2
`

type ListOffsetParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListOffset(ctx context.Context, arg ListOffsetParams) ([]Manufacturer, error) {
	rows, err := q.db.QueryContext(ctx, listOffset, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Manufacturer
	for rows.Next() {
		var i Manufacturer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.ModifiedAt,
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
