package main

import (
	"net/http"

	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/controller"
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDabase("postgres", "book_store")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "HELLO!")
	})

	r.GET("/books", controller.GetAllBooks)
	r.GET("/books/:id", controller.GetBookByID)

	r.Run()
}
