package main

import (
	"github.com/Dadamuhames/go-way/migrate"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	rootCmd.AddCommand(migrate.MigrationCommand())

	rootCmd.Execute()
}
