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

	err = createUserTable()

	if err != nil {
		return err
	}

	return nil
}

func createBookTable() error {

	query := `
    CREATE TABLE IF NOT EXISTS books (
        bookid serial unique,
        name varchar(255) not null,
        genre varchar(255) not null,
        author varchar(255) not null,
        price_inr int not null,
        count int not null,
        PRIMARY KEY (name,author)
    );
    `

	_, err := db.Exec(query)

	return err
}

func createUserTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        user_id serial,
        first_name varchar(255) not null,
        last_name varchar(255) not null,
        email varchar(255) not null unique,
        password varchar(255) not null,
        is_admin boolean not null,
        PRIMARY KEY (first_name, last_name, email)
    )
    `
	_, err := db.Exec(query)

	return err
}
