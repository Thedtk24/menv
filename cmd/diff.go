package cmd

import (
	"fmt"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/Thedtk24/menv/internal/lmod"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [name]",
	Short: "Show differences between the current environment and a saved one",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		savedEnv, err := lockfile.LoadByName(name)
		if err != nil {
			return err
		}

		currentModules := lmod.GetLoadedModules()

		savedModMap := make(map[string]bool)
		for _, m := range savedEnv.Modules {
			savedModMap[m] = true
		}
		currentModMap := make(map[string]bool)
		for _, m := range currentModules {
			currentModMap[m] = true
		}

		fmt.Printf("Differences with '%s':\n", name)

		for _, m := range savedEnv.Modules {
			if !currentModMap[m] {
				fmt.Printf("  %s %s (missing in current env)\n", GetIcon("error"), m)
			}
		}
		for _, m := range currentModules {
			if !savedModMap[m] {
				fmt.Printf("  %s %s (not in saved env)\n", GetIcon("warn"), m)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}
