package main

import (
	"github.com/Manan-Prakash-Singh/Online-BookLibrary-RestAPI/controller"
	"github.com/Manan-Prakash-Singh/Online-BookLibrary-RestAPI/middleware"
	"github.com/Manan-Prakash-Singh/Online-BookLibrary-RestAPI/models"
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
	r.POST("/books", middleware.AuthorizeAdmin, controller.CreateBook)
	r.GET("/users", middleware.AuthorizeAdmin, controller.GetAllUsers)
	r.DELETE("/books/:id", middleware.AuthorizeAdmin, controller.DeleteBook)
	r.PATCH("/user/:id/grant-admin", middleware.AuthorizeAdmin, controller.GrantAdmin)
	r.DELETE("/user/:id", middleware.AuthorizeAdmin, controller.DeleteUser)
	r.GET("/books/:id", controller.GetBookByID)
	r.POST("/user/register", controller.RegisterNewUser)
	r.POST("/user/login", controller.LoginUser)

	r.Run()
}
