package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	test1filepath = "./files/test1"
	test2filepath = "./files/test2/"
	jsonlinkpath  = "./links/json/"
	yamllinkpath  = "./links/yaml/"
	tomllinkpath  = "./links/toml/"
)

func TestReadConfig_JSON(t *testing.T) {
	path := "./testdata/test.json"
	cfg, err := readConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, cfg.Links[0].Filepath, test1filepath)
	assert.Equal(t, cfg.Links[1].Filepath, test2filepath)
	assert.Equal(t, cfg.Links[0].Linkpath, jsonlinkpath)
	assert.Equal(t, cfg.Links[1].Linkpath, jsonlinkpath)
}

func TestReadConfig_YAML(t *testing.T) {
	path := "./testdata/test.yaml"
	cfg, err := readConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, cfg.Links[0].Filepath, test1filepath)
	assert.Equal(t, cfg.Links[1].Filepath, test2filepath)
	assert.Equal(t, cfg.Links[0].Linkpath, yamllinkpath)
	assert.Equal(t, cfg.Links[1].Linkpath, yamllinkpath)
}

func TestReadConfig_TOML(t *testing.T) {
	path := "./testdata/test.toml"
	cfg, err := readConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, cfg.Links[0].Filepath, test1filepath)
	assert.Equal(t, cfg.Links[1].Filepath, test2filepath)
	assert.Equal(t, cfg.Links[0].Linkpath, tomllinkpath)
	assert.Equal(t, cfg.Links[1].Linkpath, tomllinkpath)
}

func TestReadConfig_ErrConfigNotFound(t *testing.T) {
	path := "./testdata/config.json"
	_, err := readConfig(path)
	assert.ErrorIs(t, err, ErrConfigNotFound)
}

func TestReadConfig_ErrConfigPremissionDenied(t *testing.T) {
	tmpDir := os.TempDir()
	defer os.RemoveAll(tmpDir)
	path := filepath.Join(tmpDir, "config.toml")
	_ = os.WriteFile(path, []byte(""), 0000)
	_, err := readConfig(path)
	assert.ErrorIs(t, err, ErrPermissionDenied)
}

func TestReadConfig_ErrInvalidConfig(t *testing.T) {
	path := "./testdata/incorrect_config.yaml"
	_, err := readConfig(path)
	assert.ErrorIs(t, err, ErrInvalidConfig)
}

func TestReadConfig_ErrMissingFields(t *testing.T) {
	path := "./testdata/incorrect_config.toml"
	_, err := readConfig(path)
	assert.ErrorIs(t, err, ErrMissingFields)
}
