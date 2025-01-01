package cmd

import (
	"app/src/config"
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetConfig().App.Name + " v" + config.GetConfig().App.Version)
	},
}