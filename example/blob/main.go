package main

import (
	"log"

	"github.com/xtdlib/sqlitex"
)

var schema = `
CREATE TABLE IF NOT EXISTS xxx (
    name text,
    data blob
);
`

type Data struct {
	Name string `db:"name"`
	Data []byte `db:"data"`
}

func main() {
	db := sqlitex.New("file:example.db")
	db.MustExec(schema)
	db.MustExec(`insert into xxx values ($1, $2)`, "xxx", []byte{3, 0, 4})

	var b []byte
	db.MustGet(&b, `select data from xxx limit 1`)
	log.Println("b", b)

	var dat []Data
	db.MustSelect(&dat, `select data from xxx`)

	for _, da := range dat {
		log.Printf("da: %#v", da)
	}
}
