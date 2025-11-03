package cmd

import (
	"fmt"
	"os"

	"github.com/sagar-nexturn/passwordManager/internal/repository"
	"github.com/spf13/cobra"
)

func NewRootCmd(repo repository.PasswordDbRepo) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "password-manager",
		Short: "CLI password manager",
		Long:  `Store and retrieve encrypted passwords (AES).`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the Password Manager CLI!")
		},
	}

	rootCmd.AddCommand(NewAddCmd(repo))

	return rootCmd
}

func Execute(rootCmd *cobra.Command) {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
