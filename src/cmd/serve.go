package cmd

import (
	"app/src/config"
	"app/src/connection"
	"app/src/logger"
	"app/src/models"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	logger.Info("Starting the server...")

	// Get Config
	cfg := config.GetConfig()

	// DB Connection
	connection.Connect(&cfg)
	db := connection.GetDB()
	models := append([]any{},
		&models.AuthorDetail{},
		&models.BookDetail{},
		&models.UserDetail{})
	connection.Migrate(db, models)

	time.Sleep(5 * time.Second)

	fmt.Println(connection.Ping())
	// framework := echo.New()
	// Register routes
	// pingRoutes := routes.NewPingRoutes(framework)

}
