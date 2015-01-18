package user

import (
	"github.com/m4rw3r/uuid"
	"time"
)

// User struct
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// GetId returns the user's id
func (u *User) GetId() string {
	return u.ID.String()
}
