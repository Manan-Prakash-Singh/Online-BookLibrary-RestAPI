package main

import (
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
)

func main() {
	models.InitDabase("postgres", "book_store")
}
