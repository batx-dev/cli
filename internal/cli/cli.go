package cli

import (
	"context"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Config  kong.ConfigFlag `help:"The path of configuration file." default:"~/.bat.json" type:"existingfile"`
	BaseURL string          `help:"The baseurl of batainer service." default:"https://eci.paracloud.com"`
	Token   string          `help:"The token to access batainer service."`

	Login        LoginCmd        `cmd:"" help:"Login to parauser service."`
	InstanceType InstanceTypeCmd `cmd:"" help:"List of instance types." aliases:"it"`
}

type Context struct {
	Context context.Context
	Config  kong.ConfigFlag
	BaseURL string
	Token   string
}
