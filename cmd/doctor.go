package cmd

import (
	"fmt"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/Thedtk24/menv/internal/lmod"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor [name]",
	Short: "Check if all modules of an environment are still available",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		env, err := lockfile.LoadByName(name)
		if err != nil {
			return err
		}

		hasErrors := false
		fmt.Printf("Checking environment '%s'...\n", name)
		for _, mod := range env.Modules {
			err := lmod.CheckModuleExists(mod)
			if err != nil {
				fmt.Printf("  %s %s: %v\n", GetIcon("error"), mod, err)
				hasErrors = true
			} else {
				fmt.Printf("  %s %s: ok\n", GetIcon("success"), mod)
			}
		}

		if hasErrors {
			return fmt.Errorf("some modules are missing or unavailable")
		}
		fmt.Printf("%s All modules are available.\n", GetIcon("success"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
