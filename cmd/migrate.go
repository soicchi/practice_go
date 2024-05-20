package cmd

import (
	"fmt"

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
		cfg, _ := config.NewConfig()
		db, _ := database.ConnectDB(&cfg.Database)
		database.Migrate(db)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
