package main

import (
	"strings"

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

// shortUUID shortens a UUID to the last element for display purposes.
func shortUUID(id string) string {
	a := strings.Split(id, "-")
	return a[len(a)-1]
}
