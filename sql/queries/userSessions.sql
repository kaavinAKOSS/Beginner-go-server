-- name: CreateSession :one

INSERT INTO userSessions(id,sessionId,userId)
VALUES ($1,$2,$3)
RETURNING *;

-- name: GetUserSession :one

SELECT * FROM userSessions WHERE sessionId=$1;