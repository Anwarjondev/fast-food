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

	// Create tables
	if err := createTables(); err != nil {
		log.Fatalln("Error creating tables: ", err)
	}
}

func createTables() error {
	// Create users table
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR,
			password VARCHAR,
			is_active BOOL,
			token TEXT,
			is_logged_in BOOL
		)
	`)
	if err != nil {
		return err
	}

	// Create confirm table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS confirm (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			code INT,
			is_passwed BOOL,
			is_passed BOOL
		)
	`)
	if err != nil {
		return err
	}

	// Create category table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS category (
			id SERIAL PRIMARY KEY,
			name VARCHAR
		)
	`)
	if err != nil {
		return err
	}

	// Create food table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS food (
			id SERIAL PRIMARY KEY,
			name VARCHAR,
			category_id INT REFERENCES category(id),
			img_url TEXT,
			price FLOAT,
			count_food INTEGER
		)
	`)
	if err != nil {
		return err
	}

	// Create orders table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP,
			delivered_at TIMESTAMP,
			user_id INT REFERENCES users(id),
			status VARCHAR,
			total_amount NUMERIC
		)
	`)
	if err != nil {
		return err
	}

	// Create order_detail table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS order_detail (
			id SERIAL PRIMARY KEY,
			food_id INT REFERENCES food(id),
			count INT,
			order_id INT REFERENCES orders(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
