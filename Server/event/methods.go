package event

import (
	"github.com/jmoiron/sqlx"
	"github.com/m4rw3r/uuid"
	"time"
)

type Args struct {
	ID      uuid.UUID
	Name    string
	From 	time.Time
	End 	time.Time
}

type EventMethod struct {
	DB *sqlx.DB
}


func (db *EventMethod) CreateEvent(args *Args, reply *uuid.UUID) error {
	*reply = Create(Event{Name: args.Name, From: args.From, End: args.End}, db.DB)

	return nil
}

func (db *EventMethod) GetEvent(args *Args, reply *Event) error {
	*reply = Get(args.ID, db.DB)

	return nil
}

func (db *EventMethod) DeleteEvent(args *Args, reply *error) error {
	*reply = Delete(args.ID, db.DB)

	return nil
}

func (db *EventMethod) ListEvent(args *Args, reply *[]Event) error {
	*reply = List(db.DB)

	return nil
}
