package database

import (
	"database/sql"
	"fmt"

	"github.com/Dadamuhames/go-way/internal/database/query"
)

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

func getPsqlConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		GOWAY_DIALECT,
		username,
		password,
		host,
		port,
		database,
		schema,
	)

	return sql.Open("pgx", connStr)
}

func getMysqlConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@/%s", username, password, database)
	return sql.Open("mysql", connStr)
}

func getSqlConnection(dialect string) (*sql.DB, error) {
	switch dialect {
	case "postgres":
		return getPsqlConnection()

	case "mysql":
		return getMysqlConnection()
	}

	return nil, fmt.Errorf("Unsuported dialect")
}

func getQuery(key string) string {
	return query.GetQuery(GOWAY_DIALECT, key)
}
