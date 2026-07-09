package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolve(t *testing.T) {
	hdir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		path         string
		resolvedPath string
		err          error
	}{
		{
			path:         "~/.config",
			resolvedPath: filepath.Join(hdir, ".config"),
		},
		{
			path:         "~",
			resolvedPath: hdir,
		},
		{
			path:         "",
			resolvedPath: "",
		},
		{
			path:         ".",
			resolvedPath: filepath.Join(wd, "."),
		},
		{
			path:         "..",
			resolvedPath: filepath.Join(wd, ".."),
		},
		{
			path:         "./",
			resolvedPath: filepath.Join(wd, "./"),
		},
		{
			path:         "../",
			resolvedPath: filepath.Join(wd, "../"),
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Resolve_%v", test.path), func(t *testing.T) {
			resolvedPath, err := resolve(test.path)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.resolvedPath, resolvedPath)
		})
	}
}
