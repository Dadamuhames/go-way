package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type SchemaHistory struct {
	Id       uuid.UUID
	Script   string
	Checksum int64
}

type Migration struct {
	Version     string
	Description string
	Script      string
	Checksum    int64
	Success     bool
}

func InitMigrationsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS goway_schema_history (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		version VARCHAR(50) NOT NULL,
		description VARCHAR(255) NOT NULL,
		script VARCHAR(10000) NOT NULL,
		checksum bigint NOT NULL,
		installed_on timestamp without time zone DEFAULT now(),
		success boolean not null
	); `

	err := execInTrasaction(db, query)

	if err != nil {
		fmt.Println(err)
	}
}

func FetchSchemaHistory(db *sql.DB) ([]SchemaHistory, error) {

	query := `SELECT g.id, g.script, g.checksum FROM goway_schema_history g`

	var historyEntities []SchemaHistory

	rows, err := db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("history %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var history SchemaHistory

		if err := rows.Scan(&history.Id, &history.Script, &history.Checksum); err != nil {
			return nil, fmt.Errorf("history %v", err)
		}

		historyEntities = append(historyEntities, history)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("history %v", err)
	}

	return historyEntities, err
}

func RunMigration(db *sql.DB, query string) bool {

	err := execInTrasaction(db, query)

	return err == nil
}

func StoreMigration(db *sql.DB, migration Migration) error {

	fmt.Printf("Storing migration: %v\n", migration.Version)

	query := `INSERT INTO goway_schema_history (version, description, script, checksum, success) VALUES ($1, $2, $3, $4, $5)`

	err := execInTrasaction(db, query,
		migration.Version,
		migration.Description,
		migration.Script,
		migration.Checksum,
		migration.Success,
	)

	if err != nil {
		return fmt.Errorf("SchemaHistory creation: %v", err)
	}

	return nil
}

func execInTrasaction(db *sql.DB, query string, args ...any) error {
	trx, err := db.Begin()

	if err != nil {
		return fmt.Errorf("Table init failed: %v\n", err)
	}

	defer trx.Rollback()

	_, err = trx.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %v", err)
	}

	return trx.Commit()
}
