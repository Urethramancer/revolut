package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strings"

	"github.com/Urethramancer/revolut"
)

// newClient wraps the NewClient() method in the Revolut SDK to select the configured key.
func newClient() (*revolut.Client, error) {
	key := cfg.ProductionKey
	if cfg.UseSandbox {
		key = cfg.SandboxKey
	}

	c, err := revolut.NewClient(key)
	if err != nil {
		return nil, err
	}

	c.Agent = fmt.Sprintf("Revolut Go/%s", Version[1:])
	return c, nil
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

// generateRequestID returns a string based on the last request ID in the configuration.
// We're just going for a boring, old SHA1 hash. It could easily be replaced with any hash,
// but this should be sufficient for uniqueness within one user's payments.
func generateRequestID() string {
	cfg.LastRequest++
	SaveConfig()
	id := fmt.Sprintf("%s%d", programName, cfg.LastRequest)
	h := sha1.New()
	io.WriteString(h, id)
	id = fmt.Sprintf("%x", h.Sum(nil))
	return id
}
