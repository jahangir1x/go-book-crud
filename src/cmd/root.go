package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func init() {
	RootCmd.AddCommand(ServeCmd)
	RootCmd.AddCommand(VersionCmd)
}

func Execute() {
	RootCmd.Execute()
}

