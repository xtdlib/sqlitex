package sqlitex

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/xtdlib/sqlx"
)

type DB = sqlx.DB

func Open(dataSourceName string) (*DB, error) {
	return sqlx.Open("sqlite3", dataSourceName)
}

func MustOpen(dataSourceName string) *DB {
	return sqlx.MustOpen("sqlite3", dataSourceName)
}
