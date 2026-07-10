package main

import (
	"os"
	"os/exec"
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
	test3filepath := "./testdata/files/test3"
	err := link(test3filepath, testlinkpath)
	assert.ErrorIs(t, err, nil)
	_, err = os.Lstat(filepath.Join(testlinkpath, "test3"))
	assert.ErrorIs(t, err, nil)
}

func TestUnlink(t *testing.T) {
	cPath, err := exec.LookPath("ln")
	if err != nil {
		t.Fatal(err)
	}
	test4filepath := "./testdata/files/test4"
	cmd := exec.Command(cPath, "-sf", test4filepath, filepath.Join(linkpath, "test4"))
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	err = unlink(test4filepath, linkpath)
	assert.ErrorIs(t, err, nil)
	_, err = os.Lstat(filepath.Join(linkpath, "test4"))
	assert.NotErrorIs(t, err, nil)
}

func TestUnlink_UnlinkedFile(t *testing.T) {
	test4filepath := "./testdata/files/test4"
	err := unlink(test4filepath, linkpath)
	assert.ErrorIs(t, err, nil)
}
