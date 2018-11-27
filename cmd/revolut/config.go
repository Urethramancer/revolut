package main

import (
	"os"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/slog"
)

const (
	// ConfigFile contains the main settings.
	ConfigFile = "config.json"
	// AccountsFile is the name of the account cache.
	AccountsFile = "accounts.json"
	// DetailsFile is the name of the bank details cache.
	DetailsFile = "details.json"
	// CounterpartiesFile is the name of the counterparty cache.
	CounterpartiesFile = "counterparties.json"
)

var cfg Config

// Config holds the permanent configuration.
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
	slog.Msg("Created '%s'. Edit the API keys before you run this program again.", cross.ConfigName(ConfigFile))
	os.Exit(0)
}

// SaveConfig does just that.
func SaveConfig() {
	cfgname := cross.ConfigName(ConfigFile)
	err := SaveJSON(cfgname, cfg)
	if err != nil {
		slog.Error("Error saving configuration: %s", err.Error())
		os.Exit(2)
	}
}

// LoadConfig loads the default config.
func LoadConfig() {
	cross.SetConfigPath(programName)
	cfgname := cross.ConfigName(ConfigFile)
	if !cross.FileExists(cfgname) {
		CreateConfig()
	}

	err := LoadJSON(cfgname, &cfg)
	if err != nil {
		slog.Error("Error loading configuration: %s", err.Error())
		os.Exit(2)
	}
}
