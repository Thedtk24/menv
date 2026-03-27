package cmd

import (
	"fmt"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [name]",
	Short: "Import the local .menv.lock file into saved environments",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		env, err := lockfile.LoadLocalLock()
		if err != nil {
			return err
		}

		if err := lockfile.SaveByName(name, env); err != nil {
			return fmt.Errorf("failed to save imported environment: %w", err)
		}

		fmt.Printf("%s .menv.lock imported as '%s'.\n", GetIcon("success"), name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
