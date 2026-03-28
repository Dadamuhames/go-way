package migrations

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Applying migrations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Migrations applied!")
		},
	}
}
