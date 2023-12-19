// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID        int64
	UserID    uuid.UUID
	FileName  string
	FileID    string
	CreatedAt time.Time
}

type User struct {
	ID           uuid.UUID
	UserName     string
	PasswordHash string
	CreatedAt    time.Time
}
