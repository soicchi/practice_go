package cmd

import (
	"fmt"

	"go-practice/internal/config"
	"go-practice/internal/database"
	"go-practice/internal/models/migrations"

	"github.com/spf13/cobra"
	"github.com/go-gormigrate/gormigrate/v2"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command is used to migrate the database schema to the latest version.",
	Long: `Migration is a way to keep the database schema in sync with the application's codebase.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
		cfg, _ := config.NewConfig()
		db, _ := database.ConnectDB(&cfg.Database)
		database.Migrate(db, []*gormigrate.Migration{
			migrations.UsersTable(),
		})
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
