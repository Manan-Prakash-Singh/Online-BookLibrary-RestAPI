package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/middleware"
	"github.com/Manan-Prakash-Singh/Online-Bookstore-RestAPI/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func GrantAdmin(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad URL"})
		return
	}

	user, err := models.GetUserByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = models.GrantAdmin(user.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Granted admin privileges."})
}

func GetAllUsers(c *gin.Context) {
	usersList, err := models.GetAllUsers()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		return
	}

	type customUser struct {
		UserID    int    `json:"user_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email_id"`
		IsAdmin   bool   `json:"is_admin"`
	}
	var users []customUser

	for _, user := range usersList {
		user := customUser{
			user.UserID,
			user.FirstName,
			user.LastName,
			user.Email,
			user.IsAdmin,
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad URL"})
		return
	}

	if err := models.DeleteUser(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted User"})
}
