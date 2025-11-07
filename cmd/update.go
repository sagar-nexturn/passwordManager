package cmd

import (
	"fmt"
	"time"

	"github.com/sagar-nexturn/passwordManager/internal/crypto"
	"github.com/sagar-nexturn/passwordManager/internal/repository"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(repo repository.PasswordDbRepo, cypt crypto.Crypto) *cobra.Command {
	var flagName string
	var flagOldPass string
	var flagNewPass string
	var flagUser string

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update stored password using entry name and old password",
		RunE: func(cmd *cobra.Command, args []string) error {
			if flagName == "" || flagOldPass == "" || flagNewPass == "" {
				return fmt.Errorf("name, old password and new password all are required")
			}

			p, err := repo.GetPasswordByName(flagName)
			if err != nil {
				return fmt.Errorf("failed to get password with name : %s with error: %v", flagName, err)
			}

			//Verify User
			if flagUser != p.Username {
				return fmt.Errorf("username does not match")
			}

			//Verify Password
			storedPass, err := cypt.Decrypt(p.Secret, p.Nonce)
			if err != nil {
				return err
			}
			if flagOldPass != string(storedPass) {
				return fmt.Errorf("old password does not match stored password, please try again")
			}

			ct, nonce, err := cypt.Encrypt([]byte(flagNewPass))
			if err != nil {
				return err
			}

			p.Secret = ct
			p.Nonce = nonce
			p.UpdatedAt = time.Now()

			err = repo.UpdatePassword(p)
			if err != nil {
				return fmt.Errorf("failed to update password: %v", err)
			}

			fmt.Printf("The entry with name: %v is updated:-\n", flagName)
			fmt.Printf("Updated Pass: %v\n", flagNewPass)
			return nil
		},
	}

	updateCmd.Flags().StringVarP(&flagName, "name", "n", "", "name (required)")
	updateCmd.Flags().StringVarP(&flagOldPass, "oldPass", "o", "", "Old Password (required)")
	updateCmd.Flags().StringVarP(&flagNewPass, "newPass", "p", "", "New Password (required)")
	updateCmd.Flags().StringVarP(&flagUser, "user", "u", "", "userName (optional)")
	_ = updateCmd.MarkFlagRequired("name")
	_ = updateCmd.MarkFlagRequired("oldPass")
	_ = updateCmd.MarkFlagRequired("newPass")

	return updateCmd
}
