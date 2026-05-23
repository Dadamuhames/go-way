package query

var CREATE_SCHEMA_HISTORY_TABLE_PSQL = `CREATE TABLE IF NOT EXISTS goway_schema_history (
		id VARCHER(36) PRIMARY KEY DEFAULT gen_random_uuid(),
		type VARCHAR(2) NOT NULL,
		version VARCHAR(50) NOT NULL,
		description VARCHAR(255) NOT NULL,
		script VARCHAR(10000) NOT NULL,
		checksum bigint NOT NULL,
		installed_on timestamp without time zone DEFAULT now(),
		success boolean not null
	); `

var SELECT_HISTORY_PSQL = `SELECT g.id, g.type, g.script, g.checksum FROM goway_schema_history g`

var INSERT_INTO_HISTORY_TABLE_PSQL = `INSERT INTO goway_schema_history (type, version, description, script, checksum, success) VALUES ($1, $2, $3, $4, $5, $6)`

var UPDATE_HISTORY_TABLE_PSQL = `UPDATE goway_schema_history SET checksum = $1, success = $2 WHERE script = $3`
