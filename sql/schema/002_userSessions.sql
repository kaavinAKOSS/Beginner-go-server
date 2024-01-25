-- +goose Up

CREATE TABLE userSessions(

id UUID PRIMARY KEY,
sessionId TEXT NOT NULL UNIQUE,
userId UUID REFERENCES users(id) NOT NULL UNIQUE

);

-- +goose Down

DROP TABLE userSessions;