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

// shouldDisplayCurrency checks if the given currency is in the list.
func shouldDisplayCurrency(currency, list string) bool {
	if len(list) == 0 {
		return true
	}

	l := strings.Split(list, ",")
	for _, c := range l {
		if currency == c {
			return true
		}
	}

	return false
}
