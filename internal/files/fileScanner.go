package files

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func GetScriptList() []string {

	var scripts []string

	filepath.WalkDir("./resource/db/migration/", func(path string, d fs.DirEntry, err error) error {

		fmt.Printf("Dir entity path: %s\n", path)

		if d != nil && !d.IsDir() && filepath.Ext(path) == ".sql" {
			scripts = append(scripts, path)
			fmt.Printf("File: %s\n", path)
		}

		return err
	})

	return scripts
}
