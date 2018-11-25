// Configuration editing from the command line.
package main

import (
	"errors"
	"os"

	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

// AppConfigCmd is empty.
type AppConfigCmd struct {
	SetConfig SetConfigCmd `command:"set" description:"Set configuration options."`
	GetConfig GetConfigCmd `command:"get" description:"Show configuration options."`
}

//
// Change settings.
//

// SetConfigCmd holds the commands to change settings.
type SetConfigCmd struct {
	SetProdKey SetProdKeyCmd `command:"prod" description:"Set production API key."`
	SetSandKey SetSandKeyCmd `command:"sand" description:"Set sandbox API key."`
	API        SetAPICmd     `command:"api" description:"Set the API to use."`
}

// SetProdKeyCmd changes the production API key.
type SetProdKeyCmd struct {
	Args struct {
		Key string `required:"true" positional-arg-name:"key" description:"API key to use for production."`
	} `positional-args:"true"`
}

// Execute the change.
func (cmd *SetProdKeyCmd) Execute(args []string) error {
	if !revolut.ValidKey(cmd.Args.Key) {
		return errors.New(revolut.ErrKeyFormat)
	}

	if cmd.Args.Key[0:5] != "prod_" {
		slog.Error("This is not a production key.")
		os.Exit(2)
	}

	cfg.ProductionKey = cmd.Args.Key
	SaveConfig()
	return nil
}

// SetSandKeyCmd changes the testing API key.
type SetSandKeyCmd struct {
	Args struct {
		Key string `required:"true" positional-arg-name:"key" description:"API key to use for the sandbox."`
	} `positional-args:"true"`
}

// Execute the change.
func (cmd *SetSandKeyCmd) Execute(args []string) error {
	if !revolut.ValidKey(cmd.Args.Key) {
		return errors.New(revolut.ErrKeyFormat)
	}

	if cmd.Args.Key[0:5] != "sand_" {
		slog.Error("This is not a sandbox key.")
		os.Exit(2)
	}

	cfg.SandboxKey = cmd.Args.Key
	SaveConfig()
	return nil
}

// SetAPICmd lets the user choose API to use. Default is sandbox.
type SetAPICmd struct {
	Args struct {
		API string `required:"true" positional-arg-name:"API" description:"API to use when commands are run."`
	} `positional-args:"true"`
}

// Execute the API change.
func (cmd *SetAPICmd) Execute(args []string) error {
	if len(cmd.Args.API) >= 4 {
		switch cmd.Args.API[0:4] {
		case "sand":
			slog.Msg("API set to sandbox.")
			cfg.UseSandbox = true
			SaveConfig()
			return nil
		case "prod":
			slog.Msg("API set to production.")
			cfg.UseSandbox = false
			SaveConfig()
			return nil
		}
	}

	return errors.New("unknown argument " + cmd.Args.API)
}

//
// View settings.
//

// GetConfigCmd shows configuration settings.
type GetConfigCmd struct {
	GetProdKey GetProdKeyCmd `command:"prod" description:"Show production API key."`
	GetSandKey GetSandKeyCmd `command:"sand" description:"Show sandbox API key."`
	API        GetAPICmd     `command:"api" description:"Show which API is used."`
}

// GetProdKeyCmd shows the live API key.
type GetProdKeyCmd struct{}

// Execute the key view.
func (cmd *GetProdKeyCmd) Execute(args []string) error {
	slog.Msg("%s", cfg.ProductionKey)
	return nil
}

// GetSandKeyCmd shows the test API key.
type GetSandKeyCmd struct{}

// Execute the key view.
func (cmd *GetSandKeyCmd) Execute(args []string) error {
	slog.Msg("%s", cfg.SandboxKey)
	return nil
}

// GetAPICmd shows whether sandbox or production API is being used.
type GetAPICmd struct{}

// Execute the API view.
func (cmd *GetAPICmd) Execute(args []string) error {
	if cfg.UseSandbox {
		slog.Msg("Sandbox is the active API.")
	} else {
		slog.Msg("Production is the active API.")
	}
	return nil
}
