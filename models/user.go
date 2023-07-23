package models

import (
	"database/sql"
	"fmt"
)

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

func GetAllUsers() ([]User, error) {

	rows, err := db.Query(`SELECT user_id,first_name,last_name,email,is_admin FROM users`)

	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		user := User{}
		err := rows.Scan(
			&user.UserID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.IsAdmin,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func GetUserByID(id int) (*User, error) {
	row := db.QueryRow(`SELECT * FROM users WHERE user_id = $1`, id)

	user := User{}

	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Couldn't find user in database, %s", err.Error())
		}
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {

	row := db.QueryRow(`SELECT * FROM users WHERE email = $1`, email)

	user := User{}

	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Couldn't find user in database, %s", err.Error())
		}
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GrantAdmin(id int) error {
	_, err := db.Exec(`UPDATE users SET is_admin = true WHERE user_id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
