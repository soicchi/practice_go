package cmd

import (
	"fmt"

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
		cfg, _ := config.NewConfig()
		db, _ := database.ConnectDB(&cfg.Database)
		database.RollbackLast(db)

		fmt.Println("Rollback last migration")
	},
}

func init() {
	rootCmd.AddCommand(rollbackLastCmd)
}
