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
	absPath, err := resolve(path)
	if err != nil {
		return err
	}
	absLinkPath, err := resolve(linkpath)
	if err != nil {
		return err
	}
	linkFilePath := filepath.Join(absLinkPath, filepath.Base(absPath))
	if _, err := os.Lstat(linkFilePath); err == nil {
		return nil
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
	if err := os.Symlink(absPath, linkFilePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrLinkNotFound
		}
		return ErrLinkFailed
	}
	return nil
}
