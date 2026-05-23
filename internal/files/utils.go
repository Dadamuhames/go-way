package files

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func buildMigration(path string) (error, *MigrationScript) {
	filePath := filepath.Base(path)

	version := extractVersion(filePath)
	description := extractDescription(filePath)

	return nil, &MigrationScript{
		Type:        string(filePath[0]),
		Version:     version,
		Description: description,
		Script:      path,
	}
}

func extractVersion(filename string) string {
	return strings.Split(filename, "__")[0]
}

func extractDescription(filename string) string {
	description := strings.ReplaceAll(strings.Split(filename, "__")[1], "_", " ")
	return strings.ReplaceAll(description, ".sql", "")
}

func getMigrationDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(wd, MIGRATION_DIR)
}
