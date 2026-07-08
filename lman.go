package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
)

func NewLman(stdout, stderr io.Writer, stdin io.Reader) *cli.Command {
	return &cli.Command{
		Name:                   "lman",
		Usage:                  "manage farms of symbolic links",
		Version:                "v0.0.1",
		Writer:                 stdout,
		ErrWriter:              stderr,
		Reader:                 stdin,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config.toml",
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
				Name:    "sync",
				Aliases: []string{"s"},
				Usage:   "Get all links in current folder from config",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "unlink all links in current folder",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "unlink and delete all links in current folder",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			paths := cmd.StringArgs("paths")
			if len(paths) >= 2 {
				files := paths[:len(paths)-1]
				linkfile := paths[len(paths)-1]
				for _, file := range files {
					if err := link(file, linkfile); err != nil {
						return err
					}
				}
				fmt.Fprintln(cmd.Writer, "links created")
			}
			return nil
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
