package main

import (
	"strings"

	"github.com/Urethramancer/slog"
)

type AccountCmd struct {
	AccList AccListCmd `command:"list" alias:"ls" description:"List accounts."`
}

//
// Listing accounts.
//

// AccListCmd is empty.
type AccListCmd struct {
	Short      bool   `short:"s" description:"Shorten IDs for display purposes."`
	Currencies string `short:"c" description:"List only this comma-separated list of currencies."`
}

// Execute lists the user's accounts.
func (cmd *AccListCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	list, err := c.GetAccounts()
	if err != nil {
		return err
	}

	slog.Msg("Accounts:")
	for _, acc := range list {
		id := acc.ID
		if cmd.Short {
			a := strings.Split(acc.ID, "-")
			id = a[4]
		}
		name := acc.Name
		if len(name) == 0 {
			name = "<unnamed>"
		}
		slog.Msg("%s: %s - %f %s", id, name, acc.Balance, acc.Currency)
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
