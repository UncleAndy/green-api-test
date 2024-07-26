package templates

import (
	"embed"
	"io/fs"
)

//go:embed base/*.tmpl
var BaseFS embed.FS

//go:embed content/*.tmpl
var ContentFS embed.FS

func FileNames(folder embed.FS) (files []string, err error) {
	if err := fs.WalkDir(folder, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
