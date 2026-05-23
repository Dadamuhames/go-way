package main

import (
	"fmt"
	"log"

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

			db := dbServer.GetInstance()

			defer db.Close()

			err := migrate.Migrate(db)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Migrations applied!")
		},
	}
}
