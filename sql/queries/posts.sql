-- name: CreatePost :one

INSERT INTO posts(id,title,description,authorId) VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetPostById :one

SELECT * FROM posts WHERE id=$1;

-- name: GetPostByUser :many

SELECT * FROM posts WHERE authorId=$1;