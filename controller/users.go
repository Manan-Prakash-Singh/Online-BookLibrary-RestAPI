package controller

import (
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/middleware"
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

const secretKey = "poggers69420"

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

func LoginUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.GetUserByEmail(email)

	if err != nil {
		if strings.Contains(err.Error(), "Couldn't find") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to search in database"})
			log.Printf("Unauthorized user access, email_id: %s, err: %s\n", email, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			log.Printf("Databse error: %s\n", err.Error())
		}
		c.Abort()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		log.Printf("Incorrect password: %s", err.Error())
		return
	}

	token, err := middleware.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
