package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/batx-dev/cli/parauser"
)

type LoginCmd struct {
	Username string `help:"The username to login." short:"u" required:""`
	Password string `help:"The password to login." short:"p" required:""`

	Endpoint string `help:"The endpoint to login." arg:"" default:"https://user.paratera.com"`
}

func (c *LoginCmd) Run(g *Globals) error {
	client := parauser.NewClient(c.Endpoint)
	user, err := client.LoginUser(g.Context, &parauser.LoginUserRequest{
		Email:    c.Username,
		Password: c.Password,
	})
	if err != nil {
		return fmt.Errorf("login user: %v", err)
	}

	return writeConfig(string(g.Config), map[string]any{
		"token": user.Token,
	})
}

func writeConfig(path string, kv map[string]any) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("open config file: %v", err)
	}
	defer f.Close()

	cc := make(map[string]any)
	if err := json.NewDecoder(f).Decode(&cc); err != nil {
		return fmt.Errorf("decode json config: %v", err)
	}

	for k, v := range kv {
		cc[k] = v
	}
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("truncate config file: %v", err)
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("seek config file: %v", err)
	}
	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	if err := e.Encode(cc); err != nil {
		return fmt.Errorf("encode json config: %v", err)
	}

	return nil
}
