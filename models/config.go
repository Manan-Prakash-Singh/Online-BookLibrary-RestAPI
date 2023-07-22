package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDabase(user, dbname string) {

	fmt.Println("Attemping to connect to Postgres database")

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)

	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Succesfully connected to %s.\n", dbname)

	if err := createTables(); err != nil {
		log.Fatal(err)
	}
}

func createTables() error {
	var err error

	err = createBookTable()

	if err != nil {
		return err
	}

	return nil
}

func createBookTable() error {

	query := `
    CREATE TABLE IF NOT EXISTS books (
        bookid serial primary key,
        name varchar(255) not null,
        genre varchar(255) not null,
        author varchar(255) not null,
        price_inr int,
        count int
    );
    `

	_, err := db.Exec(query)

	return err
}
