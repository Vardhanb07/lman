package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var (
	ErrLinkFailed               = errors.New("link creation failed")
	ErrLinkNotFound             = errors.New("link directory not found")
	ErrLinkFileNotFound         = errors.New("link file not found")
	ErrLinkFilePremissionDenied = errors.New("link file premission denied")
)

func link(path, linkpath string) error {
	absPath, err := homedir.Expand(path)
	if err != nil {
		return err
	}
	absLinkPath, err := homedir.Expand(linkpath)
	if err != nil {
		return err
	}
	_, err = os.Stat(absPath)
	if err != nil {
		log.Print(err)
		if errors.Is(err, os.ErrNotExist) {
			return ErrLinkFileNotFound
		}
		if errors.Is(err, os.ErrPermission) {
			return ErrLinkFilePremissionDenied
		}
		return err
	}
	if err := os.Symlink(absPath, filepath.Join(absLinkPath, filepath.Base(absPath))); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrLinkNotFound
		}
		return ErrLinkFailed
	}
	return nil
}
