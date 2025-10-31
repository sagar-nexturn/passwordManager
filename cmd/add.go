package cmd

import (
	"fmt"
	"time"

	"github.com/sagar-nexturn/passwordManager/internal/crypto"
	"github.com/sagar-nexturn/passwordManager/internal/db"
	"github.com/sagar-nexturn/passwordManager/internal/models"
	"github.com/spf13/cobra"
)

var (
	flagSite     string
	flagUsername string
	flagSecret   string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagSite == "" || flagSecret == "" {
			return fmt.Errorf("site and secret are required")
		}

		// encrypt
		ct, nonce, err := crypto.Encrypt([]byte(flagSecret))
		if err != nil {
			return err
		}

		entry := models.Password{
			Name:      flagSite,
			Username:  flagUsername,
			Secret:    ct,
			Nonce:     nonce,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Store().CreatePassword(cmd.Context(), &entry); err != nil {
			return err
		}

		fmt.Println("Added entry with id:", entry.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&flagSite, "site", "s", "", "site name (required)")
	addCmd.Flags().StringVarP(&flagUsername, "username", "u", "", "username (optional)")
	addCmd.Flags().StringVarP(&flagSecret, "secret", "p", "", "password/secret (required)")
}
