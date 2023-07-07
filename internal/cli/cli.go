package cli

import (
	"context"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Globals

	Login        LoginCmd        `cmd:"" help:"Login to parauser service."`
	InstanceType InstanceTypeCmd `cmd:"" help:"List instance types." aliases:"it"`
}

type Globals struct {
	Context context.Context `kong:"-"`

	Config  kong.ConfigFlag `help:"The path of configuration file." default:"~/.bat.json" type:"existingfile"`
	BaseURL string          `help:"The baseurl of batainer service." default:"https://eci.paracloud.com"`
	Token   string          `help:"The token to access batainer service."`
}
