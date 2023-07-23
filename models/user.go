package models

type User struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

func CreateUser(user *User) error {

	_, err := db.Exec(`INSERT INTO users(first_name,last_name,email,password,is_admin) VALUES(
        $1, $2, $3, $4, $5)`,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.IsAdmin,
	)

	if err != nil {
		return err
	}

	return nil
}
