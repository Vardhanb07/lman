package main

import (
	"os"
	"path/filepath"
)

// resolves path to abs path
// deals path starting with ~, ../, ./
func resolve(path string) (string, error) {
	hdir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if len(path) == 0 {
		return path, nil
	}
	if path[0] == '~' {
		return filepath.Join(hdir, path[1:]), nil
	}
	resolved, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return resolved, nil
}
