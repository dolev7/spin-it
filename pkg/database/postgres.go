package database

import (
	"database/sql"
	"github.com/dolev7/spin-it/pkg/logger"
	_ "github.com/lib/pq"
)

// DB is a global database connection
var DB *sql.DB

// InitDB initializes the PostgreSQL database connection
func InitDB(connStr string) error {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		logger.Log.Errorf("Error connecting to database: %v", err)
		return err
	}

	logger.Log.Debug("Connected to postgres succesfully!")

	return err
}
