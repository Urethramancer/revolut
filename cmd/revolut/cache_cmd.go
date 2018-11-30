package main

import (
	"os"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/slog"
)

// CacheCmd contains all cache manipulation.
type CacheCmd struct {
	Clear CacheClearCmd `command:"clear" description:"Clear all caches."`
}

// CacheClearCmd clears the caches.
type CacheClearCmd struct{}

// Execute the cache clearing.
func (cmd *CacheClearCmd) Execute(args []string) error {
	accounts := cross.ConfigName(AccountsFile)
	details := cross.ConfigName(DetailsFile)
	cp := cross.ConfigName(CounterpartiesFile)

	var err error
	if cross.FileExists(accounts) {
		slog.Msg("Removing %s", accounts)
		err = os.Remove(accounts)
		if err != nil {
			return err
		}
	}

	if cross.FileExists(details) {
		slog.Msg("Removing %s", details)
		err = os.Remove(details)
		if err != nil {
			return err
		}
	}

	if cross.FileExists(cp) {
		slog.Msg("Removing %s", cp)
		err = os.Remove(cp)
		if err != nil {
			return err
		}
	}

	slog.Msg("Cleared all caches.")
	return nil
}
