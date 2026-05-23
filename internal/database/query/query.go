package query

var QUERY_MAP = map[string]map[string]string{
	"postgres": {
		"CREATE": CREATE_SCHEMA_HISTORY_TABLE_PSQL,
		"INSERT": INSERT_INTO_HISTORY_TABLE_PSQL,
		"SELECT": SELECT_HISTORY_PSQL,
		"UPDATE": UPDATE_HISTORY_TABLE_PSQL,
	},
	"mysql": {
		"CREATE": CREATE_SCHEMA_HISTORY_TABLE_MYSQL,
		"INSERT": INSERT_INTO_HISTORY_TABLE_MYSQL,
		"SELECT": SELECT_HISTORY_MYSQL,
		"UPDATE": UPDATE_HISTORY_TABLE_MYSQL,
	},
}

func GetQuery(dialect string, key string) string {
	return QUERY_MAP[dialect][key]
}
