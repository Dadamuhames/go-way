package migrate

import (
	"database/sql"
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"strings"

	"github.com/Dadamuhames/go-way/internal/database"
	"github.com/Dadamuhames/go-way/internal/files"
)

func Migrate(db *sql.DB) error {

	database.InitMigrationsTable(db)

	migrationFiles := files.GetScriptList()
	storedMigrations, err := database.FetchSchemaHistory(db)

	if err != nil {
		return fmt.Errorf("Migration Error: %v", err)
	}

	fmt.Printf("Migration files: %v\n", migrationFiles)

	for _, script := range migrationFiles {
		storedMigration := findStoredByScriptPath(&storedMigrations, script)

		query, err := os.ReadFile(script)

		checksum := calculateChecksum(query)

		if err != nil {
			return fmt.Errorf("Migration Error: %v", err)
		}

		if storedMigration == nil {

			success := database.RunMigration(db, string(query))

			err, migration := buildMigration(script, checksum, success)

			if err != nil {
				return fmt.Errorf("Migration Error: %v", err)
			}

			err = database.StoreMigration(db, *migration)

			if err != nil {
				return fmt.Errorf("Migration Error: %v", err)
			}

		} else {

			if checksum != storedMigration.Checksum {
				return fmt.Errorf("Migration checksum not matching")
			}

		}
	}

	return nil
}

func findStoredByScriptPath(scripts *[]database.SchemaHistory, path string) *database.SchemaHistory {

	for _, script := range *scripts {
		if script.Script == path {
			return &script
		}
	}

	return nil
}

func buildMigration(path string, checksum int64, success bool) (error, *database.Migration) {

	version := extractVersion(path)

	description := strings.ReplaceAll(strings.Split(path, "__")[1], "_", " ")
	description = strings.ReplaceAll(description, ".sql", "")

	return nil, &database.Migration{
		Version:     version,
		Description: description,
		Script:      path,
		Checksum:    checksum,
		Success:     success,
	}
}

func extractVersion(path string) string {

	filename := filepath.Base(path)

	return strings.Split(filename, "__")[0]
}

func calculateChecksum(query []byte) int64 {
	return int64(crc32.ChecksumIEEE(query))
}
