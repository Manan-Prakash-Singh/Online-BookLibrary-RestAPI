package controller

import (
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegisterNewUser(c *gin.Context) {

	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		log.Printf("Failed to bind the response to models.User, %s\n", err)
		c.Abort()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		log.Printf("Failed to hash password, %s\n", err)
		c.Abort()
		return
	}

	user.Password = string(hashedPassword)

	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to database"})
		log.Printf("Failed to add to database, %s\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Successfully created user"})

}
