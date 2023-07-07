package cli

import (
	"context"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Config kong.ConfigFlag `help:"The path of configuration file." default:"~/.bat.json" type:"existingfile"`

	Login LoginCmd `cmd:"" help:"Login to parauser service."`
}

type Context struct {
	Context context.Context
	Config  kong.ConfigFlag
}
