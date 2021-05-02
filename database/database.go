package mockserver

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	name   string
	path   string
	schema string
	err    string
}

func (db *Database) init() {
	database, err := sql.Open("sqlite3", db.path)
	if err != nil {
		db.setError(err)
		return
	}
	defer database.Close()
	statement, err := database.Prepare(db.schema)
	if err != nil {
		db.setError(err)
		return
	}
	_, err = statement.Exec()
	if err != nil {
		db.setError(err)
	}
}

func (db *Database) setError(err error) {
	if err != nil {
		db.err = err.Error()
		log.Printf("Got error in db: %s", db.err)
	}
}

func (db *Database) Execute(query string, values ...interface{}) (*sql.Result, error) {
	database, err := sql.Open("sqlite3", db.path)
	if err != nil {
		db.setError(err)
		return nil, err
	}
	statement, err := database.Prepare(query)
	if err != nil {
		db.setError(err)
		return nil, err
	}
	result, err := statement.Exec(values...)
	if err != nil {
		db.setError(err)
		return nil, err
	}
	return &result, nil
}

func (db *Database) Select(query string) (*sql.Rows, error) {
	database, err := sql.Open("sqlite3", db.path)
	if err != nil {
		db.setError(err)
		return nil, err
	}
	rows, err := database.Query(query)
	if err != nil {
		db.setError(err)
		return nil, err
	}
	return rows, nil
}

func GetDatabase(name, path, schema string) *Database {
	db := &Database{
		name:   name,
		path:   path,
		schema: schema,
		err:    "",
	}
	db.init()
	return db
}
