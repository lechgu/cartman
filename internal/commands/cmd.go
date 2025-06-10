package commands

import (
	initialize "github.com/lechgu/cartman/internal/commands/init"
	"github.com/lechgu/cartman/internal/commands/issue"
	"github.com/lechgu/cartman/internal/commands/version"
	"github.com/lechgu/cartman/internal/meta"
	"github.com/spf13/cobra"
)

var cmd = cobra.Command{
	Use:   meta.App,
	Short: meta.Short,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	cmd.AddCommand(&version.Cmd)
	cmd.AddCommand(&initialize.Cmd)
	cmd.AddCommand(&issue.Cmd)
}
