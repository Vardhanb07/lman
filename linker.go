package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	ErrLinkFailed               = errors.New("link creation failed")
	ErrLinkNotFound             = errors.New("link directory not found")
	ErrLinkFileNotFound         = errors.New("link file not found")
	ErrLinkFilePremissionDenied = errors.New("link file premission denied")
)

// ignore linked file paths
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

// ignore unlinked file paths
func unlink(path, linkpath string) error {
	absPath, err := resolve(path)
	if err != nil {
		return err
	}
	absLinkPath, err := resolve(linkpath)
	if err != nil {
		return err
	}
	linkFilePath := filepath.Join(absLinkPath, filepath.Base(absPath))
	if _, err := os.Lstat(linkFilePath); err != nil {
		return nil
	}
	return os.Remove(linkFilePath)
}

func createLinks(files []string, linkfile string, stdout io.Writer, verbose bool) error {
	doneCh := make(chan struct{})
	errCh := make(chan error)
	filesCh := make(chan string)

	wg := sync.WaitGroup{}

	go func() {
		defer wg.Done()
		for _, file := range files {
			filesCh <- file
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Go(func() {
			for file := range filesCh {
				if verbose {
					fmt.Fprintf(stdout, "lman: creating link of %v in %v\n", file, linkfile)
				}
				if err := link(file, linkfile); err != nil {
					errCh <- err
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case <-doneCh:
			fmt.Fprintln(stdout, "lman: links created")
			return nil
		}
	}
}

func createLinksFromConfig(cfg *Config, stdout io.Writer, verbose bool) error {
	errCh := make(chan error)
	doneCh := make(chan struct{})
	linksCh := make(chan Link)

	wg := sync.WaitGroup{}

	go func() {
		defer wg.Done()
		for _, link := range cfg.Links {
			linksCh <- link
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Go(func() {
			for l := range linksCh {
				if verbose {
					fmt.Fprintf(stdout, "lman: creating link of %v in %v\n", l.Filepath, l.Linkpath)
				}
				if err := link(l.Filepath, l.Linkpath); err != nil {
					errCh <- err
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case <-doneCh:
			fmt.Fprintln(stdout, "lman: links created")
			return nil
		}
	}
}
