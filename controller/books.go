package controller

import (
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func CreateBook(c *gin.Context) {
	book := models.Book{}
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	if err := models.AddBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Book succesfully created",
	})
}

func DeleteBook(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	if err := models.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted",
	})

}
