package migrate

import (
	"database/sql"
	"fmt"
	"hash/crc32"
	"log"
	"os"

	"github.com/Dadamuhames/go-way/internal/database"
	"github.com/Dadamuhames/go-way/internal/files"
)

func findStoredByScriptPath(scripts *[]database.SchemaHistory, path string) *database.SchemaHistory {
	for _, script := range *scripts {
		if script.Script == path {
			return &script
		}
	}

	return nil
}

func mapMigration(script *files.MigrationScript, checksum int64, success bool) *database.Migration {
	return &database.Migration{
		Type:        script.Type,
		Version:     script.Version,
		Script:      script.Script,
		Checksum:    checksum,
		Description: script.Description,
		Success:     success,
	}
}

func runMigration(db *sql.DB, content []byte) bool {
	err := database.RunMigration(db, string(content))

	if err != nil {
		log.Fatal(err)
	}

	return true
}

func runVersionedMigrations(
	db *sql.DB,
	vMigrations []files.MigrationScript,
	scripts *[]database.SchemaHistory,
) error {

	var err error
	var content []byte

	for _, script := range vMigrations {
		storedMigration := findStoredByScriptPath(scripts, script.Script)

		content, err = os.ReadFile(script.Script)

		if err != nil {
			return fmt.Errorf("Migration Error: %v", err)
		}

		checksum := calculateChecksum(content)

		if storedMigration == nil {
			success := runMigration(db, content)
			migration := mapMigration(&script, checksum, success)
			err := database.StoreMigration(db, *migration)

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

func runRepeatableMigrations(
	db *sql.DB,
	rMigrations []files.MigrationScript,
	scripts *[]database.SchemaHistory,
) error {

	var err error
	var content []byte

	for _, script := range rMigrations {
		storedMigration := findStoredByScriptPath(scripts, script.Script)

		content, err = os.ReadFile(script.Script)
		checksum := calculateChecksum(content)

		if storedMigration == nil {
			success := runMigration(db, content)
			migration := mapMigration(&script, checksum, success)
			err = database.StoreMigration(db, *migration)

		} else {
			if checksum != storedMigration.Checksum {
				success := runMigration(db, content)
				migration := mapMigration(&script, checksum, success)
				err = database.UpdateMigration(db, *migration)
			}
		}
	}

	if err != nil {
		return fmt.Errorf("Migration Error: %v", err)
	}

	return nil
}

func calculateChecksum(query []byte) int64 {
	return int64(crc32.ChecksumIEEE(query))
}
