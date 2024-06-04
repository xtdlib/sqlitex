package main

import (
	"log"

	"github.com/xtdlib/sqlitex"
)

var schema = `
CREATE TABLE IF NOT EXISTS person (
    first_name text
);

CREATE TABLE IF NOT EXISTS person2 (
    first_name text,
    last_name text
);
`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	// Email     string
}

func main() {
	db := sqlitex.New("file:example.db")
	db.MustExec(schema)

	persons := []Person{}
	_ = persons
	db.MustExec(`insert into person2 values ($1, $2)`, "1", "2")
	db.MustExec(`insert into person2 values ($1, $2)`, "3", "4")
	// db.MustNamedExec(`update person2 set first_name = :first_name where last_name = :last_name`,
	// 	map[string]interface{}{
	// 		"first_name": "jane", "last_name": "doe",
	// 	})

	// row := db.QueryRowx("SELECT count(*) FROM person2").
	// if row.Err() != nil {
	// 	panic(row.Err())
	// }

	var count int
	db.MustGet(&count, "SELECT count(*) FROM person2")
	log.Println(count)

	// // err := db.Select(&persons, "SELECT * FROM person2")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// //
	// person := Person{}
	// {
	// 	rows := db.MustQueryx(`select * from person2`)
	// 	for rows.Next() {
	// 		err := rows.Scan(&person.FirstName, &person.LastName)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		log.Println(person)
	// 	}
	// 	log.Println(rows.Close())
	// }
	//
	// {
	// 	log.Println(db.MustQueryRowx(`select * from person2`).Scan(&person.FirstName, &person.LastName))
	// 	log.Println(person)
	// }

}
