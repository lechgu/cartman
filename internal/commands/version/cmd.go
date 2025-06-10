package version

import (
	"fmt"

	"github.com/lechgu/cartman/internal/meta"
	"github.com/spf13/cobra"
)

var Cmd = cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Display the %s version", meta.App),
	RunE:  version,
}

func version(cmd *cobra.Command, args []string) error {
	fmt.Println(meta.VersionString())
	if verbose && meta.Commit != "" {
		fmt.Printf("commit %s\n", meta.Commit)
	}
	return nil
}

func init() {
	Cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
