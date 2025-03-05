package users

import (
	"database/sql"
	"fmt"

	"github.com/dolev7/spin-it/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes a plain text password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with plain text
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateUser inserts a new user into the database with a hashed password
func CreateUser(email, password string) error {
	if database.PostgresDB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	_, err = database.PostgresDB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, hashedPassword)
	return err
}

// GetUserByEmail fetches a user by email
func GetUserByEmail(email string) (*User, error) {
	var user User
	if database.PostgresDB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	err := database.PostgresDB.QueryRow("SELECT id, email, password FROM users WHERE email=$1", email).Scan(
		&user.ID, &user.Email, &user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}
