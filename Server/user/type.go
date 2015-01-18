package user

import (
	"github.com/m4rw3r/uuid"
	"lab.likipe.se/go/postgresql"
	"time"
)

// User struct
type User struct {
	ID                   uuid.UUID             `json:"id"`
	Email                string                `json:"email"`
	Password             string                `json:"password,omitempty"`
	Name                 string                `json:"name"`
	Title                postgresql.NullString `json:"title"`
	SocialSecurityNumber postgresql.NullString `json:"socialSecurityNumber" db:"social_security_number"`
	Contacts             postgresql.Hstore     `json:"contacts"`
	Active               bool                  `json:"active"`
	CreatedAt            time.Time             `json:"createdAt" db:"created_at"`
}

// GetId returns the user's id
func (u *User) GetId() string {
	return u.ID.String()
}
