package dtoresponse

import "time"

type ResponseUser struct {
	ID         int64      `json:"id"`
	Role       int        `json:"role"`
	Email      string     `json:"email"`
	IsVerified bool       `json:"is_verified"`
	IsOAuth    bool       `json:"is_oauth"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
