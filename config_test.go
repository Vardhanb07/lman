package main

import (
	"testing"
)

var (
	test1filepath = "./files/test1"
	test2filepath = "./files/test2/"
	jsonlinkpath  = "./links/json/"
	yamllinkpath  = "./links/yaml/"
	tomllinkpath  = "./links/toml/"
)

func TestReadConfigWithJSON(t *testing.T) {
	path := "./testdata/test.json"
	cfg, err := readConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Links[0].Filepath != test1filepath {
		t.Errorf("expected: filepath: %v, got: filepath: %v", test1filepath, cfg.Links[0].Filepath)
	}
	if cfg.Links[1].Filepath != test2filepath {
		t.Errorf("expected: filepath: %v, got: filepath: %v", test2filepath, cfg.Links[1].Filepath)
	}
	if cfg.Links[0].Linkpath != jsonlinkpath {
		t.Errorf("expected: linkpath: %v, got: linkpath: %v", jsonlinkpath, cfg.Links[0].Linkpath)
	}
	if cfg.Links[1].Linkpath != jsonlinkpath {
		t.Errorf("expected: linkpath: %v, got: linkpath: %v", jsonlinkpath, cfg.Links[1].Linkpath)
	}
}

func TestReadConfigWithYAML(t *testing.T) {
	path := "./testdata/test.yaml"
	cfg, err := readConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Links[0].Filepath != test1filepath {
		t.Errorf("expected: filepath: %v, got: filepath: %v", test1filepath, cfg.Links[0].Filepath)
	}
	if cfg.Links[1].Filepath != test2filepath {
		t.Errorf("expected: filepath: %v, got: filepath: %v", test2filepath, cfg.Links[1].Filepath)
	}
	if cfg.Links[0].Linkpath != yamllinkpath {
		t.Errorf("expected: linkpath: %v, got: linkpath: %v", yamllinkpath, cfg.Links[0].Linkpath)
	}
	if cfg.Links[1].Linkpath != yamllinkpath {
		t.Errorf("expected: linkpath: %v, got: linkpath: %v", yamllinkpath, cfg.Links[1].Linkpath)
	}
}

func TestReadConfigWithTOML(t *testing.T) {
	path := "./testdata/test.toml"
	cfg, err := readConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Links[0].Filepath != test1filepath {
		t.Errorf("expected: filepath: %v, got: filepath: %v", test1filepath, cfg.Links[0].Filepath)
	}
	if cfg.Links[1].Filepath != test2filepath {
		t.Errorf("expected: filepath: %v, got: filepath: %v", test2filepath, cfg.Links[1].Filepath)
	}
	if cfg.Links[0].Linkpath != tomllinkpath {
		t.Errorf("expected: linkpath: %v, got: linkpath: %v", tomllinkpath, cfg.Links[0].Linkpath)
	}
	if cfg.Links[1].Linkpath != tomllinkpath {
		t.Errorf("expected: linkpath: %v, got: linkpath: %v", tomllinkpath, cfg.Links[1].Linkpath)
	}
}
