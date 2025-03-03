package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/dolev7/spin-it/docs"
	"github.com/dolev7/spin-it/internal/server"
	"github.com/dolev7/spin-it/pkg/database"
	"github.com/dolev7/spin-it/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting User Service...")

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("Error loading .env file: ", err)
	}

	// Get environment variables
	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construct database connection string
	dbConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)
	serverAddr := fmt.Sprintf("%s:%s", dbHost, port)

	// Initialize the database
	err = database.InitDB(dbConnStr)
	if err != nil {
		logger.Log.Fatalf("Database connection failed: %v", err)
	}

	// Setup the router
	router := server.SetupRouter()

	logger.Log.Info("User Service running on port ", port)

	logger.Log.Fatal(http.ListenAndServe(serverAddr, router))
}
