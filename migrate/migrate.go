package migrate

import (
	"database/sql"
	"fmt"

	"github.com/Dadamuhames/go-way/internal/database"
	"github.com/Dadamuhames/go-way/internal/files"
)

func Migrate(db *sql.DB) error {
	err := database.InitMigrationsTable(db)

	if err != nil {
		return fmt.Errorf("Schema init error: %v", err)
	}

	vMigrations, rMigrations := files.GetScriptList()

	storedMigrations, err := database.FetchSchemaHistory(db)
	if err != nil {
		return fmt.Errorf("Migration Error: %v", err)
	}

	err = runVersionedMigrations(db, vMigrations, &storedMigrations)
	if err != nil {
		return fmt.Errorf("Migration Error: %v", err)
	}

	err = runRepeatableMigrations(db, rMigrations, &storedMigrations)
	if err != nil {
		return fmt.Errorf("Migration Error: %v", err)
	}

	return nil
}
