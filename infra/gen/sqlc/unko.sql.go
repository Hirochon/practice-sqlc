// Code generated by sqlc. DO NOT EDIT.
// source: unko.sql

package db

import (
	"context"
)

const getAllUnko = `-- name: GetAllUnko :many
SELECT
       id, type, num, created_at, updated_at, deleted_at, color
FROM
    unko u
ORDER BY
    u.id DESC
`

func (q *Queries) GetAllUnko(ctx context.Context) ([]Unko, error) {
	rows, err := q.query(ctx, q.getAllUnkoStmt, getAllUnko)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Unko
	for rows.Next() {
		var i Unko
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Num,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Color,
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