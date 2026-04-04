package main

import (
	command "github.com/Dadamuhames/go-way/command/migrate"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "app"}

	rootCmd.AddCommand(command.MigrationCommand())

	rootCmd.Execute()
}
