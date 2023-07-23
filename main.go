package main

import (
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/controller"
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/middleware"
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
	r.DELETE("/books/:id", middleware.AuthorizeAdmin, controller.DeleteBook)
	r.POST("/user/register", controller.RegisterNewUser)
	r.POST("/user/login", controller.LoginUser)

	r.Run()
}
