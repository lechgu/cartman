package commands

import "github.com/spf13/cobra"

func Execute() {
	err := cmd.Execute()
	cobra.CheckErr(err)
}
