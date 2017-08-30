package pubpub

import (
	"fmt"
	"os"

	"github.com/auser/repro/cmd/db"
	"github.com/spf13/cobra"
)

var (
	// VERSION set during build
	VERSION string
)

var cfgFile string

var home = os.Getenv("HOME")

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pubpub",
	Short: "Publisher",
	Long:  `Publishing framework helper.`,
}

// Execute command
func Execute(version string) {
	VERSION = version

	db.Prepare()
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.repro.yaml)")
	RootCmd.AddCommand(cmdPrint)
	RootCmd.AddCommand(cmdRenderSass)
}
