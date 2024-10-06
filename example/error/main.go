package main

import (
	"log"

	"github.com/xtdlib/sqlitex"
)

var schema = `
CREATE TABLE IF NOT EXISTS person (
    first_name text UNIQUE,
    last_name text
);
`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	// Email     string
}

func main() {
	db := sqlitex.MustOpen("file:example.db")
	db.MustExec(schema)

	_, err := db.Exec(`insert into person values ($1, $2)`, "john", "doe")
	log.Println(sqlitex.IsErrConstraint(err))
	log.Println(sqlitex.IsErrConstraintUnique(err))

	person := Person{}
	err = db.Get(&person, `select * from person where first_name = $1`, "john3")
	log.Println(sqlitex.IsErrNoRows(err))

	// persons := []Person{}
	// err = db.Select(&persons, `select * from person where first_name = $1`, "john3")
	// log.Println(sqlitex.IsErrNoRows(err))
	// log.Println(persons)
}
