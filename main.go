package main

import (
	"app/src/cmd"
	"app/src/config"
	"app/src/logger"
)

func main() {
	// Initialize the logger
	logger.InitLogger()

	// Process the configuration file
	config.InitConfig()

	// Execute the root command
	cmd.Execute()
}
