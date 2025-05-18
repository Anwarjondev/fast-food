package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect(dsn string) {
	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Error with connecting database: ", err)
	}

}
