package manager

import (
	"bluesdp/database"

	"github.com/spf13/cobra"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Create availble data models to  Database",
		Long:  `Create data models Models to the Database. The database URI is to be provided within the migrate function or as .env variable`,
		Run: func(cmd *cobra.Command, args []string) {
			migrate()
		},
	}
)

func migrate() {

	database.MigrateDataBase()
}

func init() {
	goBlueCmd.AddCommand(migrateCmd)

}
