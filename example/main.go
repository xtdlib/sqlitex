package main

import (
	"log"

	"github.com/xtdlib/sqlitex"
)

var schema = `
CREATE TABLE IF NOT EXISTS person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE IF NOT EXISTS person2 (
    first_name text,
    last_name text
);
`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func main() {
	db := sqlitex.New("file:example.db")
	db.MustExec(schema)

	persons := []Person{}
	db.MustExec(`insert into person2 values ($1, $2)`, "john", "doe")
	db.MustNamedExec(`update person2 set first_name = :first_name where last_name = :last_name`,
		map[string]interface{}{
			"first_name": "jane", "last_name": "doe",
		})
	err := db.Select(&persons, "SELECT * FROM person2")
	if err != nil {
		panic(err)
	}

	log.Println(persons)
}
