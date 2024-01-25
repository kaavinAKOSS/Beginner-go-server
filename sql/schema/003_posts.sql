-- +goose Up

CREATE TABLE posts(
id UUID PRIMARY KEY,
title TEXT NOT NULL,
description TEXT NOT NULL,
authorId UUID REFERENCES users(id) NOT NULL
);

-- +goose Down

DROP TABLE posts;
