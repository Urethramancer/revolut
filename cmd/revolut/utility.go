package main

import (
	"github.com/Urethramancer/revolut"
)

// newClient wraps the NewClient() method in the Revolut SDK to select the configured key.
func newClient() (*revolut.Client, error) {
	key := cfg.ProductionKey
	if cfg.UseSandbox {
		key = cfg.SandboxKey
	}

	return revolut.NewClient(key)
}
