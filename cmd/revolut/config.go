package main

import (
	"os"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/slog"
)

const (
	// ConfigName contains the main settings.
	ConfigName = "config.json"
	// AccountsName is the name of the account cache.
	AccountsName = "accounts.json"
)

var cfg Config

type Config struct {
	// ProductionKey has access to the production API where changes actually matter.
	ProductionKey string `json:"production_key"`
	// SandboxKey is for testing and experimenting.
	SandboxKey string `json:"sandbox_key"`
	UseSandbox bool   `json:"usesandbox"`
}

// CreateConfig creates a default configuration file which will need the API keys changed.
func CreateConfig() {
	cfg.ProductionKey = "change me"
	cfg.SandboxKey = "change me"
	cfg.UseSandbox = true
	SaveConfig()
	slog.Msg("Created '%s'. Edit the API keys before you run this program again.", cross.ConfigName(ConfigName))
	os.Exit(0)
}

// SaveConfig does just that.
func SaveConfig() {
	cfgname := cross.ConfigName(ConfigName)
	err := SaveJSON(cfgname, cfg)
	if err != nil {
		slog.Error("Error saving configuration: %s", err.Error())
		os.Exit(2)
	}
}

// LoadConfig loads the default config.
func LoadConfig() {
	cross.SetConfigPath(programName)
	cfgname := cross.ConfigName(ConfigName)
	if !cross.FileExists(cfgname) {
		CreateConfig()
	}

	err := LoadJSON(cfgname, &cfg)
	if err != nil {
		slog.Error("Error loading configuration: %s", err.Error())
		os.Exit(2)
	}
}
