package cmd

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command is used to migrate the database schema to the latest version.",
	Long:  "If you want to migrate the database schema to the latest version, you can use this command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")

		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		db, err := database.ConnectDB(&cfg.Database)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err := database.Migrate(db); err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}

		log.Println("Migrate database schema")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
