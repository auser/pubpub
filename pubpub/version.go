package pubpub

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(RootCmd.Use + " " + VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
