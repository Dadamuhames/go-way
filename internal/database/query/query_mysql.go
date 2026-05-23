package query

var CREATE_SCHEMA_HISTORY_TABLE_MYSQL = `CREATE TABLE IF NOT EXISTS goway_schema_history (
    id          VARCHAR(36)   PRIMARY KEY DEFAULT (UUID()),
    type        VARCHAR(2)    NOT NULL,
    version     VARCHAR(50)   NOT NULL,
    description VARCHAR(255)  NOT NULL,
    script      VARCHAR(10000) NOT NULL,
    checksum    BIGINT        NOT NULL,
    installed_on DATETIME     DEFAULT CURRENT_TIMESTAMP,
    success     BOOLEAN       NOT NULL
);`

var SELECT_HISTORY_MYSQL = `SELECT g.id, g.type, g.script, g.checksum FROM goway_schema_history g`

var INSERT_INTO_HISTORY_TABLE_MYSQL = `INSERT INTO goway_schema_history (type, version, description, script, checksum, success) VALUES (?, ?, ?, ?, ?, ?)`

var UPDATE_HISTORY_TABLE_MYSQL = `UPDATE goway_schema_history SET checksum = ?, success = ? WHERE script = ?`
