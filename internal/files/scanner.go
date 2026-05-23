package files

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

func GetScriptList() ([]MigrationScript, []MigrationScript) {
	var versionedScripts []MigrationScript
	var repeatableScripts []MigrationScript

	r, _ := regexp.Compile(`^[VR][0-9]*__.*\.sql$`)

	migrationsRelativeDir := fmt.Sprintf("./%s", MIGRATION_DIR)

	filepath.WalkDir(migrationsRelativeDir, func(path string, d fs.DirEntry, err error) error {
		filePath := strings.Replace(path, MIGRATION_DIR, "", 1)

		if d != nil && !d.IsDir() && r.MatchString(filePath) {
			err, script := buildMigration(path)

			if err != nil {
				return err
			}

			if script.Type == "V" {
				versionedScripts = append(versionedScripts, *script)
			} else {
				repeatableScripts = append(repeatableScripts, *script)
			}
		}

		return err
	})

	return versionedScripts, repeatableScripts
}
