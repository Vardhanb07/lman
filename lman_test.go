package main_test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
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

func TestLman_WithPathsVerbose(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	err := lman.Run(context.Background(), []string{"lman", "--verbose", test1filepath, test2filepath, linkpath})
	if err != nil {
		t.Fatal(err)
	}
	out := stdout.String()
	assert.Contains(t, out, "creating link")
	assert.Contains(t, out, "links created")
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test1filepath))); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test2filepath))); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLman_WithConfigExtJSON(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	jsonConfigFile := "./testdata/test.json"
	err := lman.Run(context.Background(), []string{"lman", "--config", jsonConfigFile})
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

func TestLman_WithConfigExtJSONVerbose(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	jsonConfigFile := "./testdata/test.json"
	err := lman.Run(context.Background(), []string{"lman", "--config", jsonConfigFile, "--verbose"})
	if err != nil {
		t.Fatal(err)
	}
	out := stdout.String()
	assert.Contains(t, out, "creating link")
	assert.Contains(t, out, "links created")
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test1filepath))); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test2filepath))); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLman_WithConfigExtTOML(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	tomlConfigFile := "./testdata/test.toml"
	err := lman.Run(context.Background(), []string{"lman", "--config", tomlConfigFile})
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

func TestLman_WithConfigExtTOMLVerbose(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	tomlConfigFile := "./testdata/test.toml"
	err := lman.Run(context.Background(), []string{"lman", "--config", tomlConfigFile, "--verbose"})
	if err != nil {
		t.Fatal(err)
	}
	out := stdout.String()
	assert.Contains(t, out, "creating link")
	assert.Contains(t, out, "links created")
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test1filepath))); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test2filepath))); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLman_WithConfigExtYAML(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	yamlConfigFile := "./testdata/test.yaml"
	err := lman.Run(context.Background(), []string{"lman", "--config", yamlConfigFile})
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

func TestLman_WithConfigExtYAMLVerbose(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	yamlConfigFile := "./testdata/test.toml"
	err := lman.Run(context.Background(), []string{"lman", "--config", yamlConfigFile})
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

func TestLman_DefaultConfig(t *testing.T) {
	t.Chdir("./testdata/")
	lman := main.NewLman(stdout, stderr, stdin)
	err := lman.Run(context.Background(), []string{"lman"})
	if err != nil {
		t.Fatal(err)
	}
	t.Chdir("../")
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

func TestLman_DefaultConfigVerbose(t *testing.T) {
	t.Chdir("./testdata/")
	lman := main.NewLman(stdout, stderr, stdin)
	err := lman.Run(context.Background(), []string{"lman", "--verbose"})
	if err != nil {
		t.Fatal(err)
	}
	t.Chdir("../")
	out := stdout.String()
	assert.Contains(t, out, "creating link")
	assert.Contains(t, out, "links created")
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test1filepath))); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(linkpath, filepath.Base(test2filepath))); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLman_ConfigExtUnsupported(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	err := lman.Run(context.Background(), []string{"lman", "--config", "./testdata/test.conf"})
	assert.ErrorIs(t, err, main.ErrUnsupportedConfigFileFormat)
}

func TestLman_DefaultConfigNotFound(t *testing.T) {
	lman := main.NewLman(stdout, stderr, stdin)
	err := lman.Run(context.Background(), []string{"lman", "--verbose"})
	assert.ErrorIs(t, err, main.ErrDefaultConfigFileNotFound)
}

func TestLman_Remove(t *testing.T) {
	cPath, err := exec.LookPath("bash")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(cPath, "setup_test_remove.sh")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	lman := main.NewLman(stdout, stderr, stdin)
	t.Chdir("./test/")
	if err := lman.Run(context.Background(), []string{"lman", "remove"}); err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat("lman")
	assert.Contains(t, err.Error(), "no such file or directory")
	_, err = os.Stat("lman.config.toml")
	assert.Equal(t, err, nil)
	t.Chdir("../")
	t.Cleanup(func() {
		if err := os.RemoveAll("./test/"); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLman_RemoveVerbose(t *testing.T) {
	cPath, err := exec.LookPath("bash")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(cPath, "setup_test_remove.sh")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	lman := main.NewLman(stdout, stderr, stdin)
	t.Chdir("./test/")
	if err := lman.Run(context.Background(), []string{"lman", "remove", "--verbose"}); err != nil {
		t.Fatal(err)
	}
	out := stdout.String()
	assert.Contains(t, out, "deleting")
	t.Chdir("../")
	t.Cleanup(func() {
		if err := os.RemoveAll("./test/"); err != nil {
			t.Fatal(err)
		}
	})
}
