package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Manifest map[string]string

func manifestPath(root string) string {
	return filepath.Join(root, ".guard.json")
}

func WriteManifest(root string, m Manifest) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(manifestPath(root), data, 0644)
}

func ReadManifest(root string) (Manifest, error) {
	data, err := os.ReadFile(manifestPath(root))
	if err != nil {
		return nil, err
	}

	var m Manifest
	err = json.Unmarshal(data, &m)
	return m, err
}