package user

import (
	"github.com/m4rw3r/uuid"
	"lab.likipe.se/go/postgresql"
)

func Create(u User, db postgresql.Ext) uuid.UUID {
	id := uuid.UUID{}
	row := db.QueryRowx(`INSERT INTO "Users" ("email", "password") VALUES ($1, $2) RETURNING "id"`,
		u.Email, u.Password)

	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

func List(db postgresql.Ext) []User {
	u := make([]User, 0)

	err := db.Select(&u, `SELECT "id", "email", "password", "created_at"
	FROM "Users"`)
	if err != nil {
		panic(err)
	}

	return u

}

func GetProfile(ID uuid.UUID, db postgresql.Ext) User {
	u := User{}

	err := db.Get(&u, `SELECT "id", "email", "password", "created_at"
	FROM "Users"
	WHERE  id = $1  
	`, ID)
	if err != nil {
		panic(err)
	}

	return u
}
