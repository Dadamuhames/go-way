package files

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
)

func GetScriptList() ([]MigrationScript, []MigrationScript) {
	var versionedScripts []MigrationScript
	var repeatableScripts []MigrationScript

	r, _ := regexp.Compile(`^[VR][0-9]*__.*\.sql$`)

	migrationsRelativeDir := getMigrationDir()

	filepath.WalkDir(migrationsRelativeDir, func(path string, d fs.DirEntry, err error) error {

		fmt.Printf("Discovered file: %s\n", path)

		filePath := filepath.Base(path)

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
