package main

import (
	"os"
	"strings"

	"github.com/Urethramancer/cross"
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

	details := loadAccounts()

	slog.Msg("Accounts:")
	for _, acc := range list {
		id := acc.ID
		if !details.HasID(id) {
			det, err := c.GetAccountDetails(id)
			if err != nil {
				return err
			}
			details.Add(id, det)
		}
		if cmd.Short {
			a := strings.Split(acc.ID, "-")
			id = a[4]
		}
		name := acc.Name
		if len(name) == 0 {
			name = "<unnamed>"
		}
		slog.Msg("%s (%s): %s - %f %s", id, acc.State, name, acc.Balance, acc.Currency)
	}

	saveAccounts(details)
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

// loadAccounts loads the cached account details.
func loadAccounts() *DetailsMap {
	fn := cross.ConfigName(AccountsName)
	det := DetailsMap{}
	err := LoadJSON(fn, &det)
	if err != nil {
		slog.Warn("Couldn't load accounts cache: %s. Proceeding with clean slate.", err.Error())
	}

	return &det
}

// saveAccounts saves the account details cache.
func saveAccounts(det *DetailsMap) {
	fn := cross.ConfigName(AccountsName)
	err := SaveJSON(fn, det)
	if err != nil {
		slog.Error("Error saving account details: ", err.Error())
		os.Exit(2)
	}
}
