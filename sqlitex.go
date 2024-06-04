package sqlitex

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sqlx.DB
}

func New(dbname string) *DB {
	db, err := sqlx.Connect("sqlite", dbname)
	if err != nil {
		panic(err)
	}
	return &DB{
		DB: db,
	}
}

func (db *DB) MustNamedExec(query string, arg interface{}) {
	_, err := db.NamedExec(query, arg)
	if err != nil {
		panic(err)
	}
}
