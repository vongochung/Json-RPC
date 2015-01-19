package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/m4rw3r/uuid"
	"github.com/pengux/check"
	"errors"
)

type Args struct {
	ID       uuid.UUID
	Email    string
	Password string
}

type UserMethod struct {
	DB *sqlx.DB
}

var (
	validate         = check.Email{}
	validatePassword = check.MinChar{8}
)

func (db *UserMethod) CreateUser(args *Args, reply *uuid.UUID) error {
	// Validate
	errs := validate.Validate(args.Email)
	if errs != nil {
		return errors.New("Email is invalid")
	}

	_,e := FindOneByEmail(db.DB, args.Email)
	if e == nil {
		return errors.New("Email has existed")
	}

	errs = validatePassword.Validate(args.Password)
	if errs != nil {
		return errors.New("Password length is greater 8")
	}


	*reply = Create(User{Email: args.Email, Password: args.Password}, db.DB)

	return nil
}

func (db *UserMethod) GetUser(args *Args, reply *User) error {
	*reply = Get(args.ID, db.DB)

	return nil
}

func (db *UserMethod) DeleteUser(args *Args, reply *error) error {
	*reply = Delete(args.ID, db.DB)

	return nil
}

func (db *UserMethod) ListUser(args *Args, reply *[]User) error {
	*reply = List(db.DB)

	return nil
}
