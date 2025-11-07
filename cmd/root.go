package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/sagar-nexturn/passwordManager/internal/crypto"
	"github.com/sagar-nexturn/passwordManager/internal/repository"
	"github.com/spf13/cobra"
)

func NewRootCmd(repo repository.PasswordDbRepo, cypt crypto.Crypto) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "password-manager",
		Short: "CLI password manager",
		Long:  `Store and retrieve encrypted passwords (AES).`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the Password Manager CLI!")
		},
	}

	rootCmd.AddCommand(NewAddCmd(repo, cypt))
	rootCmd.AddCommand(NewGetCmd(repo, cypt))
	rootCmd.AddCommand(NewDeleteCmd(repo))
	rootCmd.AddCommand(NewUpdateCmd(repo, cypt))

	return rootCmd
}

func Execute(rootCmd *cobra.Command) {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
