package models

type User struct {
	UserID    int    `json:"user_id"`
	FirstName int    `json:"first_name"`
	LastName  int    `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}
