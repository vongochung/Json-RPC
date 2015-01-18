package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/m4rw3r/uuid"
)

type Args struct {
	ID uuid.UUID
}

type UserMethod struct {
	DB *sqlx.DB
}

func (db *UserMethod) CreateUser(args *Args, reply *uuid.UUID) error {
	*reply = Create(User{Email: "holy.vo@likipe.se", Password: "123456"}, db.DB)

	return nil
}

func (db *UserMethod) Profile(args *Args, reply *User) error {
	*reply = GetProfile(args.ID, db.DB)

	return nil
}
