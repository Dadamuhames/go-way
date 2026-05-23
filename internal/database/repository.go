package database

import (
	"database/sql"
	"fmt"
)

func InitMigrationsTable(db *sql.DB) error {
	return execInTrasaction(db, getQuery("CREATE"))
}

func FetchSchemaHistory(db *sql.DB) ([]SchemaHistory, error) {
	var historyEntities []SchemaHistory

	rows, err := db.Query(getQuery("SELECT"))

	if err != nil {
		return nil, fmt.Errorf("history %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var history SchemaHistory

		if err := rows.Scan(&history.Id, &history.Type, &history.Script, &history.Checksum); err != nil {
			return nil, fmt.Errorf("history %v", err)
		}

		historyEntities = append(historyEntities, history)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("history %v", err)
	}

	return historyEntities, err
}

func RunMigration(db *sql.DB, query string) error {
	return execInTrasaction(db, query)
}

func StoreMigration(db *sql.DB, migration Migration) error {
	return execInTrasaction(
		db,
		getQuery("INSERT"),
		migration.Type,
		migration.Version,
		migration.Description,
		migration.Script,
		migration.Checksum,
		migration.Success,
	)
}

func UpdateMigration(db *sql.DB, migration Migration) error {
	return execInTrasaction(
		db,
		getQuery("UPDATE"),
		migration.Checksum,
		migration.Success,
		migration.Script,
	)
}
