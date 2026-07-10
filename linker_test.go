package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testlinkpath    = "./testdata/links"
	testfilepath    = "./testdata/files/test1"
	testfiledirpath = "./testdata/files/test2/"
)

func TestLink(t *testing.T) {
	err := link(testfilepath, testlinkpath)
	assert.ErrorIs(t, err, nil)
	_, err = os.Lstat(filepath.Join(testlinkpath, filepath.Base(testfilepath)))
	assert.NotErrorIs(t, err, os.ErrNotExist)
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(testlinkpath, filepath.Base(testfilepath))); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLink_ErrLinkFileNotFound(t *testing.T) {
	tmpfile := "does-not-exist"
	err := link(tmpfile, testlinkpath)
	assert.ErrorIs(t, err, ErrLinkFileNotFound)
}

func TestLink_ErrLinkNotFound(t *testing.T) {
	tmpfile := "does-not-exist"
	err := link(testfilepath, tmpfile)
	assert.ErrorIs(t, err, ErrLinkNotFound)
}

func TestLink_WithDir(t *testing.T) {
	err := link(testfiledirpath, testlinkpath)
	assert.ErrorIs(t, err, nil)
	_, err = os.Lstat(filepath.Join(testlinkpath, filepath.Base(testfiledirpath)))
	assert.NotErrorIs(t, err, os.ErrNotExist)
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(testlinkpath, filepath.Base(testfiledirpath))); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLink_LinkedFile(t *testing.T) {
}

func TestUnlink(t *testing.T) {
}

func TestUnlink_WithDir(t *testing.T) {
}

func TestUnlink_UnlinkedFile(t *testing.T) {
}
