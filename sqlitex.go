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

func (db *DB) MustGet(dest interface{}, query string, args ...interface{}) {
	err := db.Get(dest, query, args...)
	if err != nil {
		panic(err)
	}
}

func (db *DB) MustSelect(dest interface{}, query string, args ...interface{}) {
	err := db.Select(dest, query, args...)
	if err != nil {
		panic(err)
	}
}

func (db *DB) MustQueryx(query string, args ...interface{}) *sqlx.Rows {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		panic(err)
	}
	return rows
}

func (db *DB) MustQueryRowx(query string, args ...interface{}) *sqlx.Row {
	row := db.QueryRowx(query, args...)
	if row.Err() != nil {
		panic(row.Err())
	}
	return row
}
