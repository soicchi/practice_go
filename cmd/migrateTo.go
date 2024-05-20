package cmd

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"

	"github.com/spf13/cobra"
)

var migrateToCmd = &cobra.Command{
	Use:   "migrateTo",
	Short: "Migrate to a specific migration",
	Long:  `Given a migration ID, migrateTo will apply all migrations up to and including the migration with the provided ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateTo called")

		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		db, err := database.ConnectDB(&cfg.Database)
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		if err := database.MigrateTo(db, args[0]); err != nil {
			log.Fatalf("failed to migrate to migration: %v", err)
		}

		log.Printf("Migrate to version: %s", args[0])
	},
}

func init() {
	rootCmd.AddCommand(migrateToCmd)
}
