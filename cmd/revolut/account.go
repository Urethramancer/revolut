package main

import (
	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

type AccountCmd struct {
	AccList AccListCmd `command:"list" alias:"ls" description:"List accounts."`
}

//
// Listing accounts.
//

// AccListCmd is empty.
type AccListCmd struct{}

// Execute lists the user's accounts.
func (cmd *AccListCmd) Execute(args []string) error {
	c, err := revolut.NewClient(cfg.SandboxKey)
	if err != nil {
		return err
	}

	list, err := c.GetAccounts()
	if err != nil {
		return err
	}

	slog.Msg("Accounts:")
	for _, acc := range list {
		slog.Msg("%s: %s - %f %s", acc.ID, acc.Name, acc.Balance, acc.Currency)
	}

	return nil
}

//
// Account caching.
//

// AccUpdateCmd fetches a list of accounts and stores the details for each locally for faster lookup.
type AccUpdateCmd struct{}

func (cmd *AccUpdateCmd) Execute(args []string) error {
	return nil
}
