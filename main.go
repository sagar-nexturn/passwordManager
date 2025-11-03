package main

import (
	"log"

	"github.com/sagar-nexturn/passwordManager/cmd"
	"github.com/sagar-nexturn/passwordManager/internal/config"
	"github.com/sagar-nexturn/passwordManager/internal/repository"
)

func main() {
	//Loading .env file for credentials
	config.LoadEnv()

	//Connecting to PostgresSQL
	db := config.InitDB()
	log.Println("Connected to database on render successfully!!!")

	// Initialize Repository
	passwordRepo := repository.NewPostgresPasswordRepo(db)

	// Build root command with repo injected
	rootCmd := cmd.NewRootCmd(passwordRepo)

	// Execute CLI
	cmd.Execute(rootCmd)
}
