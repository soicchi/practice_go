package cmd

import (
	"fmt"
	"log"

	"go-practice/internal/config"
	"go-practice/internal/database"

	"github.com/spf13/cobra"
)

// rollbackLastCmd represents the rollbackLast command
var rollbackLastCmd = &cobra.Command{
	Use:   "rollbackLast",
	Short: "Rollback the last migration.",
	Long:  "If you want to rollback the last migration, you can use this command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rollbackLast called")

		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		db, err := database.ConnectDB(&cfg.Database)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err := database.RollbackLast(db); err != nil {
			log.Fatalf("failed to rollback last migration: %v", err)
		}

		log.Println("Rollback last migration")
	},
}

func init() {
	rootCmd.AddCommand(rollbackLastCmd)
}
