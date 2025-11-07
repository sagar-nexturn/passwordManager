package main

import (
	"log"
	"os"

	"github.com/sagar-nexturn/passwordManager/cmd"
	"github.com/sagar-nexturn/passwordManager/internal/config"
	"github.com/sagar-nexturn/passwordManager/internal/crypto"
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

	// Initialize and choose Crypto Logic
	var cypt crypto.Crypto
	if os.Getenv("USE_KMS") == "true" {
		cypt = crypto.NewKMSCrypto()
		log.Println("Using AWS KMS for encryption")
	} else {
		cypt = crypto.NewAESCrypto()
		log.Println("Using local AES encryption")
	}

	// Build root command with repo and crypto logic injected
	rootCmd := cmd.NewRootCmd(passwordRepo, cypt)

	// Execute CLI
	cmd.Execute(rootCmd)
}
