package entity

import (
	"time"
)

type User struct {
	ID         int64
	Role       int
	Email      string
	IsVerified bool
	IsOAuth    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
