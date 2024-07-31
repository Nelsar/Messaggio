package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	config "messaggio.com/configuration"
)

func DBConnect() (*sql.DB, error) {
	c := config.GetConfiguration()

	db, err := sql.Open("postgres", c.DSN)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return db, err
}
