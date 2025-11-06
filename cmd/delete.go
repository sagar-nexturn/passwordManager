package cmd

import (
	"fmt"

	"github.com/sagar-nexturn/passwordManager/internal/repository"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(repo repository.PasswordDbRepo) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Remove stored passwords",
	}

	deleteCmd.AddCommand(NewDeleteByIdCmd(repo))
	deleteCmd.AddCommand(NewDeleteByNameCmd(repo))

	return deleteCmd
}

func NewDeleteByIdCmd(repo repository.PasswordDbRepo) *cobra.Command {
	var flagId string

	deleteByIdCmd := &cobra.Command{
		Use:   "byId",
		Short: "Remove stored password using entry id",
		RunE: func(cmd *cobra.Command, args []string) error {
			if flagId == "" {
				return fmt.Errorf("id is required")
			}

			err := repo.DeletePasswordById(flagId)
			if err != nil {
				return fmt.Errorf("failed to delete password: %v", err)
			}

			fmt.Printf("The entry with Id: %v is deleted", flagId)
			return nil
		},
	}

	deleteByIdCmd.Flags().StringVarP(&flagId, "id", "i", "", "id (required)")
	_ = deleteByIdCmd.MarkFlagRequired("id")

	return deleteByIdCmd
}

func NewDeleteByNameCmd(repo repository.PasswordDbRepo) *cobra.Command {
	var flagName string

	getByNameCmd := &cobra.Command{
		Use:   "byName",
		Short: "Remove stored password using entry name",
		RunE: func(cmd *cobra.Command, args []string) error {
			if flagName == "" {
				return fmt.Errorf("name is required")
			}

			err := repo.DeletePasswordByName(flagName)
			if err != nil {
				return fmt.Errorf("failed to get password: %v", err)
			}

			fmt.Printf("The entry with name: %v is deleted", flagName)
			return nil
		},
	}

	getByNameCmd.Flags().StringVarP(&flagName, "name", "n", "", "name (required)")
	_ = getByNameCmd.MarkFlagRequired("name")

	return getByNameCmd
}
