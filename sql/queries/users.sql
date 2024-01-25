-- name: SignUpUser :one

INSERT INTO users(id,name,password,created_at)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetUserByID :one

SELECT * FROM users WHERE id=$1;

-- name: GetExistingUser :one

SELECT * FROM users WHERE name=$1 AND password=$2;
