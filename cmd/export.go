package cmd

import (
	"fmt"
	"strings"

	"github.com/Thedtk24/menv/internal/lockfile"
	"github.com/spf13/cobra"
)

var currentDir bool

var exportCmd = &cobra.Command{
	Use:   "export [name]",
	Short: "Export an environment to .menv.lock, or load it if --current-dir is used",
	RunE: func(cmd *cobra.Command, args []string) error {
		if currentDir {
			env, err := lockfile.LoadLocalLock()
			if err != nil {
				return nil
			}
			if len(env.Modules) > 0 {
				fmt.Printf("module purge;\nmodule load %s;\n", strings.Join(env.Modules, " "))
			}
			return nil
		}

		if len(args) != 1 {
			return fmt.Errorf("an environment name is required unless --current-dir is used")
		}

		name := args[0]
		env, err := lockfile.LoadByName(name)
		if err != nil {
			return err
		}

		if err := lockfile.SaveLocalLock(env); err != nil {
			return err
		}

		fmt.Printf("%s Environment '%s' exported to .menv.lock.\n", GetIcon("success"), name)
		return nil
	},
}

func init() {
	exportCmd.Flags().BoolVar(&currentDir, "current-dir", false, "Generate module load commands from local .menv.lock")
	rootCmd.AddCommand(exportCmd)
}
