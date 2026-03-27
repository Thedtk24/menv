package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [name]",
	Short: "Edit an environment's YAML file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not find user home directory: %w", err)
		}
		path := filepath.Join(home, ".menv", fmt.Sprintf("%s.yaml", name))

		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("environment '%s' not found", name)
		}

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}

		c := exec.Command(editor, path)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			return fmt.Errorf("failed to open editor: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
