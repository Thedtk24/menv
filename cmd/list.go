package cmd

import (
	"fmt"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		envs, err := lockfile.ListAll()
		if err != nil {
			return fmt.Errorf("failed to list environments: %w", err)
		}

		if len(envs) == 0 {
			fmt.Println("No saved environments found.")
			return nil
		}

		fmt.Println("Available environments:")
		for _, env := range envs {
			fmt.Printf("  %s %s\n", GetIcon("success"), env)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
