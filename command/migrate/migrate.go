package command

import (
	"fmt"

	"github.com/Dadamuhames/go-way/internal/database"
	"github.com/Dadamuhames/go-way/migrate"
	"github.com/spf13/cobra"
)

func MigrationCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Applying migrations",
		Run: func(cmd *cobra.Command, args []string) {

			dbServer := database.New()

			err := migrate.Migrate(dbServer.GetInstance())

			if err != nil {
				fmt.Printf("Migration error: %v", err)
				return
			}

			fmt.Println("Migrations applied!")
		},
	}
}
