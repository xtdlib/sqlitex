package main

import (
	"database/sql/driver"
	"encoding/json"
	"log"
	"os"

	"github.com/xtdlib/sqlitex"
)

var schema = `'
CREATE TABLE IF NOT EXISTS queue (
    command TEXT
	env json
);
`

type StringSlice []string

func (t *StringSlice) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), t)
}

func (t *StringSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(t)
	return string(b), err
}

type command struct {
	Command string
	Env     StringSlice
}

func main() {
	os.Remove("example.db")
	db := sqlitex.MustOpen("file:example.db")
	db.MustExec(`create table if not exists queue ( command text, env json )`)
	db.MustExec(`insert into queue (command, env) values (?, ?)`, "echo hello", `["FOO=bar", "BAZ=qux"]`)
	queue := &command{}
	// env := &StringSlice{}
	// env := ""
	db.Get(queue, `select * from queue limit 1`)
	// log.Printf("%#+v", env)
	log.Printf("%#+v", queue)
	// db.MustExec(`insert into queue (command, env) values (?, ?)`, "echo hello", `xcxx`)
}
