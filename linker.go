package main

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrLinkFailed               = errors.New("link creation failed")
	ErrLinkNotFound             = errors.New("link directory not found")
	ErrLinkFileNotFound         = errors.New("link file not found")
	ErrLinkFilePremissionDenied = errors.New("link file premission denied")
)

func link(path, linkpath string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	_, err = os.Stat(absPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrLinkFileNotFound
		}
		if errors.Is(err, os.ErrPermission) {
			return ErrLinkFilePremissionDenied
		}
		return err
	}
	if err := os.Symlink(absPath, filepath.Join(linkpath, filepath.Base(absPath))); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrLinkNotFound
		}
		return ErrLinkFailed
	}
	return nil
}
