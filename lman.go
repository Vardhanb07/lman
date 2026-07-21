package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

	"github.com/urfave/cli/v3"
)

var (
	ErrUnsupportedConfigFileFormat = errors.New("config file format not supported")
	ErrDefaultConfigFileNotFound   = errors.New("default config file not found, default files are 'lman.config.toml', 'lman.config.yaml', 'lman.config.json'. If you have custom config file provide it with --config, -c flag")
)

var (
	version = "dev"
)

func defaultConfigFiles() []string {
	return []string{"lman.config.toml", "lman.config.yaml", "lman.config.json"}
}

func defaultConfigExts() []string {
	return []string{".toml", ".yaml", ".json"}
}

func NewLman(stdout, stderr io.Writer, stdin io.Reader) *cli.Command {
	return &cli.Command{
		Name:                   "lman",
		Usage:                  "manage farms of symbolic links",
		Version:                version,
		Writer:                 stdout,
		ErrWriter:              stderr,
		Reader:                 stdin,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Set the config file path",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"V"},
				Value:   false,
				Usage:   "Use this for verbose output",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArgs{
				Name: "paths",
				Min:  0,
				Max:  -1,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "unlink",
				Aliases: []string{"u"},
				Usage:   "unlink all soft links defined in config file",
				Arguments: []cli.Argument{
					&cli.StringArgs{
						Name: "links",
						Min:  0,
						Max:  -1,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					links := cmd.StringArgs("links")
					switch {
					case len(links) != 0:
						return removeLinks(links)
					default:
						cfgFiles := defaultConfigFiles()
						wd, err := os.Getwd()
						if err != nil {
							return err
						}
						var cfgFile string
						for _, file := range cfgFiles {
							fstat, err := os.Stat(filepath.Join(wd, file))
							if os.IsNotExist(err) {
								continue
							} else if err != nil {
								return err
							}
							cfgFile = fstat.Name()
							break
						}
						if cfgFile == "" {
							cfgFile = cmd.String("config")
							if cfgFile == "" {
								return ErrDefaultConfigFileNotFound
							}
						}
						cfg, err := readConfig(cfgFile)
						return removeLinksFromConfig(cfg)
					}
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			paths := cmd.StringArgs("paths")
			verbose := cmd.Bool("verbose")
			cfgFile := cmd.String("config")
			switch {
			case len(paths) >= 2:
				files := paths[:len(paths)-1]
				linkfile := paths[len(paths)-1]
				return createLinks(files, linkfile, cmd.Writer, verbose)
			case cfgFile != "":
				exts := defaultConfigExts()
				if !slices.Contains(exts, filepath.Ext(cfgFile)) {
					return ErrUnsupportedConfigFileFormat
				}
				if verbose {
					fmt.Fprintf(cmd.Writer, "lman: reading config file\n")
				}
				cfg, err := readConfig(cfgFile)
				if err != nil {
					return err
				}
				return createLinksFromConfig(cfg, cmd.Writer, verbose)
			default:
				cfgFiles := defaultConfigFiles()
				wd, err := os.Getwd()
				if err != nil {
					return err
				}
				var cfgFile string
				for _, file := range cfgFiles {
					fstat, err := os.Stat(filepath.Join(wd, file))
					if os.IsNotExist(err) {
						continue
					} else if err != nil {
						return err
					}
					cfgFile = fstat.Name()
					break
				}
				if cfgFile == "" {
					return ErrDefaultConfigFileNotFound
				}
				cfg, err := readConfig(cfgFile)
				return createLinksFromConfig(cfg, cmd.Writer, verbose)
			}
		},
	}
}

func main() {
	lman := NewLman(os.Stdout, os.Stderr, os.Stdin)
	if err := lman.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "lman: %v\n", err)
		os.Exit(1)
	}
}
