package users

import (
	"github.com/dolev7/spin-it/pkg/database"
)

func CreateUser(email, password string) error {
	_, err := database.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)
	return err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := database.DB.QueryRow("SELECT id, email, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}
