package main

import (
	"fmt"
	"path/filepath"
)

func Init(root string) error {
	files, err := ScanFiles(root)
	if err != nil {
		return err
	}

	manifest := make(Manifest)

	for _, file := range files {
		hash, err := HashFile(file)
		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(root, file)
		manifest[rel] = hash
	}

	return WriteManifest(root, manifest)
}

func Verify(root string) error {
	stored, err := ReadManifest(root)
	if err != nil {
		return err
	}

	currentFiles, err := ScanFiles(root)
	if err != nil {
		return err
	}

	current := make(map[string]string)

	for _, file := range currentFiles {
		hash, err := HashFile(file)
		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(root, file)
		current[rel] = hash
	}

	for path, oldHash := range stored {
		newHash, ok := current[path]
		if !ok {
			fmt.Println(path, "REMOVED")
			continue
		}

		if oldHash == newHash {
			fmt.Println(path, "OK")
		} else {
			fmt.Println(path, "FAIL")
		}
	}

	for path := range current {
		if _, ok := stored[path]; !ok {
			fmt.Println(path, "ADDED")
		}
	}

	return nil
}