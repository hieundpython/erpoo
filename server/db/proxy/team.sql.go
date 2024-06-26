// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: team.sql

package dbproxy

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams
(id, "name", description)
VALUES($1, $2, $3)
RETURNING id, name, description
`

type CreateTeamParams struct {
	ID          pgtype.UUID `json:"id"`
	Name        pgtype.Text `json:"name"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRow(ctx, createTeam, arg.ID, arg.Name, arg.Description)
	var i Team
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const deleteTeam = `-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id=$1
`

func (q *Queries) DeleteTeam(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteTeam, id)
	return err
}

const getListTeams = `-- name: GetListTeams :many
SELECT id, name, description FROM teams
ORDER BY name
`

func (q *Queries) GetListTeams(ctx context.Context) ([]Team, error) {
	rows, err := q.db.Query(ctx, getListTeams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeam = `-- name: GetTeam :one
SELECT id, name, description FROM teams
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTeam(ctx context.Context, id pgtype.UUID) (Team, error) {
	row := q.db.QueryRow(ctx, getTeam, id)
	var i Team
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const updateTeam = `-- name: UpdateTeam :one
UPDATE teams
SET "name"=$1, description=$2
WHERE id=$3
RETURNING id, name, description
`

type UpdateTeamParams struct {
	Name        pgtype.Text `json:"name"`
	Description pgtype.Text `json:"description"`
	ID          pgtype.UUID `json:"id"`
}

func (q *Queries) UpdateTeam(ctx context.Context, arg UpdateTeamParams) (Team, error) {
	row := q.db.QueryRow(ctx, updateTeam, arg.Name, arg.Description, arg.ID)
	var i Team
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}
