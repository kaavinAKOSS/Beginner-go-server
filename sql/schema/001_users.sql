-- +goose Up

CREATE TABLE users(

id UUID PRIMARY KEY,
name TEXT NOT NULL,
password TEXT NOT NULL,
created_at TIMESTAMP NOT NULL

);

-- +goose Down
DROP TABLE users;
