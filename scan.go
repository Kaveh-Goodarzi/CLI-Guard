package main

import (
	"os"
	"path/filepath"
)

func ScanFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Name() == ".guard.json" {
			return nil
		}

		if d.Type().IsRegular() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}