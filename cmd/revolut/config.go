package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/slog"
)

const (
	ConfigName = "config.json"
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
	var data []byte
	var err error
	data, err = json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		slog.Error("Error: %s", err.Error())
		os.Exit(2)
	}

	cfgname := cross.ConfigName(ConfigName)
	err = ioutil.WriteFile(cfgname, data, 0600)
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

	data, err := ioutil.ReadFile(cfgname)
	if err != nil {
		slog.Error("Error loading configuration: %s", err.Error())
		os.Exit(2)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		slog.Error("Error: %s\n", err.Error())
		os.Exit(2)
	}
}
