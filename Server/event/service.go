package event

import (
	"github.com/m4rw3r/uuid"
	"lab.likipe.se/go/postgresql"
)

func Create(e Event, db postgresql.Ext) uuid.UUID {
	id := uuid.UUID{}
	row := db.QueryRowx(`INSERT INTO "Events" ("name", "from", "end") VALUES ($1,$2,$3) RETURNING "id"`,
		e.Name, e.From,e.End)

	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

func List(db postgresql.Ext) []Event {
	e := make([]Event, 0)

	err := db.Select(&e, `SELECT "id", "name", "from", "end"
	FROM "Events"`)
	if err != nil {
		panic(err)
	}

	return e

}

func Get(ID uuid.UUID, db postgresql.Ext) Event {
	e := Event{}

	err := db.Get(&e, `SELECT "id", "name", "from", "end"
	FROM "Events"
	WHERE  id = $1  
	`, ID)
	if err != nil {
		panic(err)
	}

	return e
}

func Delete(ID uuid.UUID, db postgresql.Ext) error {
	_, err := db.Exec(`DELETE FROM "Events" WHERE id = $1 `, ID)
	if err != nil {
		panic(err)
	}

	return err
}

