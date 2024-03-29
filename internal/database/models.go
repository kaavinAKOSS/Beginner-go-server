// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID
	Title       string
	Description string
	Authorid    uuid.UUID
}

type User struct {
	ID        uuid.UUID
	Name      string
	Password  string
	CreatedAt time.Time
}

type Usersession struct {
	ID        uuid.UUID
	Sessionid string
	Userid    uuid.UUID
}
