package cmd

import (
	"fmt"
	"log"

	"github.com/sagar-nexturn/passwordManager/internal/crypto"
	"github.com/sagar-nexturn/passwordManager/internal/repository"
	"github.com/spf13/cobra"
)

func NewGetCmd(repo repository.PasswordDbRepo) *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve and decrypt stored passwords",
	}

	getCmd.AddCommand(NewGetByIdCmd(repo))
	getCmd.AddCommand(NewGetByNameCmd(repo))
	getCmd.AddCommand(NewGetAllCmd(repo))

	return getCmd
}

func NewGetByIdCmd(repo repository.PasswordDbRepo) *cobra.Command {
	var flagId string

	getByIdCmd := &cobra.Command{
		Use:   "byId",
		Short: "Retrieve and decrypt stored password using entry id",
		RunE: func(cmd *cobra.Command, args []string) error {
			if flagId == "" {
				return fmt.Errorf("id is required")
			}

			p, err := repo.GetPasswordByID(flagId)
			if err != nil {
				return fmt.Errorf("failed to get password: %v", err)
			}

			pt, err := crypto.Decrypt(p.Secret, p.Nonce)
			if err != nil {
				return err
			}

			fmt.Printf("The entry for Id: %v is Name: %s and password: %s", flagId, p.Name, pt)
			return nil
		},
	}

	getByIdCmd.Flags().StringVarP(&flagId, "id", "i", "", "id (required)")
	_ = getByIdCmd.MarkFlagRequired("id")

	return getByIdCmd
}

func NewGetByNameCmd(repo repository.PasswordDbRepo) *cobra.Command {
	var flagName string

	getByNameCmd := &cobra.Command{
		Use:   "byName",
		Short: "Retrieve and decrypt stored password using entry name",
		RunE: func(cmd *cobra.Command, args []string) error {
			if flagName == "" {
				return fmt.Errorf("name is required")
			}

			p, err := repo.GetPasswordByName(flagName)
			if err != nil {
				return fmt.Errorf("failed to get password: %v", err)
			}

			pt, err := crypto.Decrypt(p.Secret, p.Nonce)
			if err != nil {
				return err
			}

			fmt.Printf("The entry for Name: %s is password: %s", p.Name, pt)
			return nil
		},
	}

	getByNameCmd.Flags().StringVarP(&flagName, "name", "n", "", "name (required)")
	_ = getByNameCmd.MarkFlagRequired("name")

	return getByNameCmd
}

func NewGetAllCmd(repo repository.PasswordDbRepo) *cobra.Command {
	getAllCmd := &cobra.Command{
		Use:   "all",
		Short: "Retrieve and decrypt all stored passwords",
		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := repo.GetAllPasswords()
			if err != nil {
				return fmt.Errorf("failed to get passwords: %v", err)
			}

			passwords := make(map[string][]byte)

			for _, v := range p {
				pt, err := crypto.Decrypt(v.Secret, v.Nonce)
				if err != nil {
					log.Printf("failed to decrypt password for name: %s, with error info: %v", v.Name, err)
					continue
				}
				passwords[v.Name] = pt
			}

			fmt.Println("The name-passwords pair are:")
			for name, password := range passwords {
				fmt.Printf("%s -> %s\n", name, password)
			}
			return nil
		},
	}

	return getAllCmd
}
