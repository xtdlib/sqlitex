package sqlitex

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"

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

func IsErrBusy(err error) bool {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code == sqlite3.ErrBusy {
			return true
		}
	}
	return false
}

func IsErrConstraint(err error) bool {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			return true
		}
	}
	return false
}

func IsErrConstraintUnique(err error) bool {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return true
		}
	}
	return false
}

func IsErrNoRows(err error) bool {
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}
	return false
}
