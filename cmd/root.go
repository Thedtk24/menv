package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

func SetVersion(v string) {
	version = v
	rootCmd.Version = v
}

var rootCmd = &cobra.Command{
	Use:     "menv",
	Short:   "menv - Lmod/Tcl module environment manager",
	Version: version,
	Long:    "menv is a CLI tool to easily save, load, and manage HPC module lists.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s Error: %v\n", GetIcon("error"), err)
		os.Exit(1)
	}
}

func GetIcon(icon string) string {
	if os.Getenv("NO_COLOR") != "" {
		return ""
	}
	switch icon {
	case "success":
		return "✓"
	case "error":
		return "✗"
	case "warn":
		return "⚠"
	case "download":
		return "📥"
	}
	return icon
}
