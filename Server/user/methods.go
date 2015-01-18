package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/m4rw3r/uuid"
)

type Args struct {
	ID       uuid.UUID
	Email    string
	Password string
}

type UserMethod struct {
	DB *sqlx.DB
}

func (db *UserMethod) CreateUser(args *Args, reply *uuid.UUID) error {
	*reply = Create(User{Email: args.Email, Password: args.Password}, db.DB)

	return nil
}

func (db *UserMethod) Profile(args *Args, reply *User) error {
	*reply = GetProfile(args.ID, db.DB)

	return nil
}

func (db *UserMethod) ListUser(args *Args, reply *[]User) error {
	*reply = List(db.DB)
	return nil
}
