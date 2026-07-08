package main_test

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	main "github.com/vardhanb07/lman"
)

var (
	stdout        = &bytes.Buffer{}
	stderr        = &bytes.Buffer{}
	stdin         = strings.NewReader("")
	test1filepath = "./testdata/files/test1"
	test2filepath = "./testdata/files/test2"
	linkpath      = "./testdata/links"
)

func TestLman_WithPaths(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	err := lman.Run(context.Background(), []string{"lman", test1filepath, test2filepath, linkpath})
	if err != nil {
		t.Fatal(err)
	}
	out := stdout.String()
	assert.Contains(t, out, "links created")
	_, err = os.Lstat(filepath.Join(linkpath, filepath.Base(test1filepath)))
	assert.ErrorIs(t, err, nil)
	_, err = os.Lstat(filepath.Join(linkpath, filepath.Base(test2filepath)))
	assert.ErrorIs(t, err, nil)
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test1filepath))); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test2filepath))); err != nil {
			t.Fatal(err)
		}
	})
}
