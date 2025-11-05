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

	//Creating tables and inserting sample data
	//PostgresInit := repository.NewPostgresInitRepo(db)
	//err := PostgresInit.CreatePasswordsTableIfNotExist()
	//if err != nil {
	//	log.Fatalf("Error creating passwords table: %v", err)
	//}
	//log.Println("Table 'passwords' ensured in database.")

	//err = PostgresInit.InsertSampleData()
	//if err != nil {
	//	log.Fatalf("Error inserting sample data: %v", err)
	//}
	//log.Println("Sample data ensured in database.")

	// Initialize Repository
	passwordRepo := repository.NewPostgresPasswordRepo(db)

	// Build root command with repo injected
	rootCmd := cmd.NewRootCmd(passwordRepo)

	// Execute CLI
	cmd.Execute(rootCmd)
}
