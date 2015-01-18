package bin

import (
	"lab.likipe.se/go/postgresql"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func DB() *sqlx.DB {
	if db == nil {
		db = postgresql.CreateConnection(Config)
	}

	return db
}
