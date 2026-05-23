package files

import (
	"path/filepath"
	"strings"
)

func buildMigration(path string) (error, *MigrationScript) {
	version := extractVersion(path)
	description := extractDescription(path)

	filePath := strings.Replace(path, MIGRATION_DIR, "", 1)

	return nil, &MigrationScript{
		Type:        string(filePath[0]),
		Version:     version,
		Description: description,
		Script:      path,
	}
}

func extractVersion(path string) string {
	filename := filepath.Base(path)
	return strings.Split(filename, "__")[0]
}

func extractDescription(path string) string {
	description := strings.ReplaceAll(strings.Split(path, "__")[1], "_", " ")
	return strings.ReplaceAll(description, ".sql", "")
}
