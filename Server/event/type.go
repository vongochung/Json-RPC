package event

import (
	"github.com/m4rw3r/uuid"
	"time"
)

// Event struct
type Event struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	From  	time.Time `json:"from"`
	End 	time.Time `json:"end"`
}

// GetId returns the user's id
func (e *Event) GetId() string {
	return e.ID.String()
}
