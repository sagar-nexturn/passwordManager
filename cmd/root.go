package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "password-manager",
	Short: "CLI password manager",
	Long:  `Store and retrieve encrypted passwords (AES).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Password Manager CLI!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (optional)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}
