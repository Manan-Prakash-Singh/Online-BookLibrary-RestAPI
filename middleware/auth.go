package middleware

import (
	"fmt"
	"github.com/Manan-Prakash-Singh/Online-BookLibrary-RestAPI/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const secretKey = "poggers69420"

func GenerateToken(user *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.UserID,
		"email":    user.Email,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
	})

	return token.SignedString([]byte(secretKey))
}

func AuthorizeUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		err := verifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthorizeAdmin(c *gin.Context) {

	err := verifyToken(c)

	if err != nil {
		log.Printf("Error in verifying the token, %s\n", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	admin, _ := c.Get("is_admin")

	if !admin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Insufficient Permission",
		})
		c.Abort()
		return
	}

	c.Next()
}

func verifyToken(c *gin.Context) error {

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		return fmt.Errorf("Missing Authorization header")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("Invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Println("Couldn't get the claims")
		return fmt.Errorf("Invalid JWT")
	}

	userFloatID, ok := claims["id"].(float64)

	if !ok {
		log.Println("Error extracting user_id")
		return fmt.Errorf("Invalid JWT")
	}

	email, ok := claims["email"].(string)

	if !ok {
		log.Println("Error in extracting the email")
		return fmt.Errorf("Invalid JWT")
	}

	admin, ok := claims["is_admin"].(bool)

	if !ok {
		log.Println("Error in extracting admin field")
		return fmt.Errorf("Invalid JWT")
	}

	userID := int(userFloatID)

	c.Set("email", email)
	c.Set("is_admin", admin)
	c.Set("user_id", userID)

	return nil
}
