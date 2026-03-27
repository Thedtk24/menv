package cmd

import (
	"fmt"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename [old_name] [new_name]",
	Short: "Rename an environment",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		oldName := args[0]
		newName := args[1]

		env, err := lockfile.LoadByName(oldName)
		if err != nil {
			return err
		}

		if err := lockfile.SaveByName(newName, env); err != nil {
			return fmt.Errorf("failed to save environment with new name: %w", err)
		}

		if err := lockfile.DeleteByName(oldName); err != nil {
			return fmt.Errorf("new environment created, but failed to delete old environment: %w", err)
		}

		fmt.Printf("%s Environment '%s' renamed to '%s'.\n", GetIcon("success"), oldName, newName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
