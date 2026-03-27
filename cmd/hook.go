package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Generate shell (Bash/Zsh) integration",
	Run: func(cmd *cobra.Command, args []string) {
		script := `
_menv_hook() {
    if [ -f ".menv.lock" ]; then
        eval "$(menv export --current-dir)"
    fi
}

# Zsh integration
if [ -n "$ZSH_VERSION" ]; then
    autoload -Uz add-zsh-hook
    add-zsh-hook chpwd _menv_hook
# Bash integration
elif [ -n "$BASH_VERSION" ]; then
    PROMPT_COMMAND="_menv_hook; $PROMPT_COMMAND"
fi
`
		fmt.Println(script)
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
