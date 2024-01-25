// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: userSessions.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one

INSERT INTO userSessions(id,sessionId,userId)
VALUES ($1,$2,$3)
RETURNING id, sessionid, userid
`

type CreateSessionParams struct {
	ID        uuid.UUID
	Sessionid string
	Userid    uuid.UUID
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Usersession, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.ID, arg.Sessionid, arg.Userid)
	var i Usersession
	err := row.Scan(&i.ID, &i.Sessionid, &i.Userid)
	return i, err
}

const getUserSession = `-- name: GetUserSession :one

SELECT id, sessionid, userid FROM userSessions WHERE sessionId=$1
`

func (q *Queries) GetUserSession(ctx context.Context, sessionid string) (Usersession, error) {
	row := q.db.QueryRowContext(ctx, getUserSession, sessionid)
	var i Usersession
	err := row.Scan(&i.ID, &i.Sessionid, &i.Userid)
	return i, err
}
