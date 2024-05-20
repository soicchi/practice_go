package cmd

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"

	"github.com/spf13/cobra"
)

var rollbackToCmd = &cobra.Command{
	Use:   "rollbackTo",
	Short: "Rollback to a specific migration",
	Long:  `Given a migration ID, rollbackTo will revert all migrations up to and including the migration with the provided ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rollbackTo called")

		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		db, err := database.ConnectDB(&cfg.Database)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err := database.RollbackTo(db, args[0]); err != nil {
			log.Fatalf("failed to rollback to migration: %v", err)
		}

		log.Println("Rollback to migration")
	},
}

func init() {
	rootCmd.AddCommand(rollbackToCmd)
}
