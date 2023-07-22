package main

import (
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/controller"
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	models.InitDabase("postgres", "book_store")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "HELLO!")
	})

	r.GET("/books", controller.GetAllBooks)
	r.POST("/books", controller.CreateBook)
	r.GET("/books/:id", controller.GetBookByID)

	r.Run()
}
