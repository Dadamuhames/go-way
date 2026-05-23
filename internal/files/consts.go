package files

import "fmt"

var (
	MIGRATION_DIR               = "resource/db/migration/"
	RELATIVE_MIGRATION_DIR_PATH = fmt.Sprintf("./%s", MIGRATION_DIR)
)
