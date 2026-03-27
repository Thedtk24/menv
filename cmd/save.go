package cmd

import (
	"fmt"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/Thedtk24/menv/internal/lmod"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [name]",
	Short: "Save the current module environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		modules := lmod.GetLoadedModules()

		env := &lockfile.Environment{
			Modules: modules,
			EnvVars: make(map[string]string),
		}

		if err := lockfile.SaveByName(name, env); err != nil {
			return fmt.Errorf("failed to save environment: %w", err)
		}

		fmt.Printf("%s Environment '%s' successfully saved (%d modules).\n", GetIcon("success"), name, len(modules))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
