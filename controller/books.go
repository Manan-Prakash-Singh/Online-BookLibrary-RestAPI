package controller

import (
	"net/http"
	"strconv"

	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := models.GetBookByID(id)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, book)
}
