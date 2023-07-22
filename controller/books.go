package controller

import (
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindBooks(c *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, books)
}
