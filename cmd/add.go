package cmd

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sagar-nexturn/passwordManager/internal/crypto"
	"github.com/sagar-nexturn/passwordManager/internal/models"
	"github.com/sagar-nexturn/passwordManager/internal/repository"
	"github.com/spf13/cobra"
)

func NewAddCmd(repo repository.PasswordDbRepo) *cobra.Command {
	var (
		flagSite     string
		flagUsername string
		flagSecret   string
	)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new password entry",
		RunE: func(cmd *cobra.Command, args []string) error {
			if flagSite == "" || flagSecret == "" {
				return fmt.Errorf("site and secret are required")
			}

			// Encrypt password before saving
			ct, nonce, err := crypto.Encrypt([]byte(flagSecret))
			if err != nil {
				return err
			}

			entry := &models.Password{
				ID:        uuid.New().String(),
				Name:      flagSite,
				Username:  flagUsername,
				Secret:    ct,
				Nonce:     nonce,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := repo.AddPassword(entry); err != nil {
				return fmt.Errorf("failed to add password: %v", err)
			}

			fmt.Println("✅ Added entry with ID:", entry.ID)
			return nil
		},
	}

	addCmd.Flags().StringVarP(&flagSite, "site", "s", "", "site name (required)")
	addCmd.Flags().StringVarP(&flagUsername, "username", "u", "", "username (optional)")
	addCmd.Flags().StringVarP(&flagSecret, "secret", "p", "", "password/secret (required)")
	_ = addCmd.MarkFlagRequired("site")
	_ = addCmd.MarkFlagRequired("secret")

	return addCmd
}
