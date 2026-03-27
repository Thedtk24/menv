package cmd

import (
	"fmt"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a saved environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := lockfile.DeleteByName(name); err != nil {
			return fmt.Errorf("failed to delete environment: %w", err)
		}
		fmt.Printf("%s Environment '%s' deleted.\n", GetIcon("success"), name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
