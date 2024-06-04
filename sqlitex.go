package sqlitex

import (
	"database/sql"

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

func (db *DB) MustNamedExec(query string, arg interface{}) sql.Result {
	return must1(db.NamedExec(query, arg))
}

func (db *DB) MustGet(dest interface{}, query string, args ...interface{}) {
	must(db.Get(dest, query, args...))
}

func (db *DB) MustSelect(dest interface{}, query string, args ...interface{}) {
	must(db.Select(dest, query, args...))
}

func (db *DB) MustQueryx(query string, args ...interface{}) *sqlx.Rows {
	return must1(db.Queryx(query, args...))
}

// NOTES(aca): not sure this is required
func (db *DB) MustQueryRowx(query string, args ...interface{}) *sqlx.Row {
	row := db.QueryRowx(query, args...)
	if row.Err() != nil {
		panic(row.Err())
	}
	return row
}

func must(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func must1[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func must2[T1 any, T2 any](v T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return v, v2
}

func must3[T1 any, T2 any, T3 any](v T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return v, v2, v3
}
