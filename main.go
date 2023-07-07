package main

import (
	"context"

	"github.com/alecthomas/kong"

	"github.com/batx-dev/cli/internal/cli"
)

func main() {
	ctx := kong.Parse(&cli.CLI, kong.Configuration(kong.JSON))
	err := ctx.Run(&cli.Context{Context: context.Background(), Config: cli.CLI.Config})
	ctx.FatalIfErrorf(err)
}
