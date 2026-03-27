package cmd

import (
	"fmt"
	"strings"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/spf13/cobra"
)

var loadScript bool

var loadCmd = &cobra.Command{
	Use:   "load [name]",
	Short: "Load a module environment (or output shell commands via --script)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		env, err := lockfile.LoadByName(name)
		if err != nil {
			return err
		}

		if loadScript {
			if len(env.Modules) > 0 {
				fmt.Printf("module purge;\nmodule load %s;\n", strings.Join(env.Modules, " "))
			} else {
				fmt.Printf("module purge;\n")
			}
			return nil
		}

		fmt.Printf("%s To load '%s', run:\n", GetIcon("warn"), name)
		fmt.Printf("eval \"$(menv load %s --script)\"\n", name)
		return nil
	},
}

func init() {
	loadCmd.Flags().BoolVar(&loadScript, "script", false, "Emit shell commands (module load ...)")
	rootCmd.AddCommand(loadCmd)
}
