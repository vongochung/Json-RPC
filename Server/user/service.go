package user

import (
	"github.com/m4rw3r/uuid"
	"lab.likipe.se/go/postgresql"
	"strings"
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

func Get(ID uuid.UUID, db postgresql.Ext) User {
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

func Delete(ID uuid.UUID, db postgresql.Ext) error {
	_, err := db.Exec(`DELETE FROM "Users" WHERE id = $1 `, ID)
	if err != nil {
		panic(err)
	}

	return err
}

// FindOneByEmail fetch a specific user by query
func FindOneByEmail(db postgresql.Ext, email string) (*User, error) {
	u := User{}

	err := db.Get(&u, `SELECT
		u."id", u."email", u."password", u."created_at"
	FROM
		"Users" u
	WHERE
		email = $1
	LIMIT 1`, strings.ToLower(email))

	return &u, err
}

