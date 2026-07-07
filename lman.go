package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func run() error {
	cmd := &cli.Command{
		Name:    "lman",
		Usage:   "manage farms of symbolic links",
		Version: "v0.0.1",
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
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}
	return cmd.Run(context.Background(), os.Args)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
