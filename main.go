package main

import (
	"context"

	"github.com/alecthomas/kong"

	"github.com/batx-dev/cli/internal/cli"
)

func main() {
	c := cli.CLI{
		Globals: cli.Globals{
			Context: context.Background(),
		},
	}
	ctx := kong.Parse(&c,
		kong.Name("bat"),
		kong.Description("A command-line tool for managing batainer service."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Configuration(kong.JSON),
	)
	err := ctx.Run(&c.Globals)
	ctx.FatalIfErrorf(err)
}
