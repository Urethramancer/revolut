package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

// CounterpartyCmd holds tool commands for viewing and modifying counterparties.
type CounterpartyCmd struct {
	List   CPListCmd   `command:"list" alias:"ls" description:"List counterparties. Will fetch from cache if available."`
	Update CPUpdateCmd `command:"update" alias:"up" description:"Update counterparty cache."`
	Get    CPGetCmd    `command:"get" description:"Get a counterparty by nickname or UUID. Will fetch from cache if available."`
	Add    CPAddCmd    `command:"add" description:"Add a new counterparty."`
	Delete CPDeleteCmd `command:"delete" alias:"del" alias:"rm" description:"Delete a counterparty."`
}

// CPListCmd shows a list of counterparties.
type CPListCmd struct {
	DefaultShowOptions
}

// Execute the counterparty list command.
func (cmd *CPListCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	cache := &CounterpartyCache{}
	cache.Load()
	if len(*cache) == 0 {
		list, err := c.GetCounterparties()
		if err != nil {
			return err
		}

		for _, cp := range list {
			cache.Set(cp.ID, cp)
		}
		cache.Save()
	}

	if len(*cache) == 0 {
		slog.Msg("No counterparties to list.")
		return nil
	}

	n := len(*cache)
	if n == 1 {
		slog.Msg("1 counterparty:")
	} else {
		slog.Msg("%d counterparties:", len(*cache))
	}

	for _, cp := range *cache {
		displayCounterparty(&cp, cmd.Short, cmd.Details)
	}
	return nil
}

func displayCounterparty(cp *revolut.Counterparty, short, details bool) {
	id := cp.ID
	if short {
		id = shortUUID(id)
	}

	slog.Msg("%s (%s): %s (%s), Phone: %s, updated %s", id, cp.Type, cp.Name, cp.Country, cp.Phone, cp.UpdatedAt.Format(time.RFC822))
	if !details || len(cp.Accounts) == 0 {
		return
	}

	slog.Msg("\tBank details:")
	for _, acc := range cp.Accounts {
		id = acc.ID
		if short {
			id = shortUUID(id)
		}
		slog.Msg("\t%s (%s, %s)", id, acc.Type, acc.Currency)
		if acc.Type == "external" {
			slog.Msg("\t\t%s", acc.Name)
			if len(acc.Account) > 0 {
				slog.Msg("\t\tAccount no.: %s", acc.Account)
			}
			if len(acc.SortCode) > 0 {
				slog.Msg("\t\tSort code: %s", acc.SortCode)
			}
			if len(acc.Email) > 0 {
				slog.Msg("\t\tE-mail: %s", acc.Email)
			}
			if len(acc.Country) > 0 {
				slog.Msg("\t\tBank country: %s", acc.Country)
			}
			if len(acc.Charges) > 0 {
				slog.Msg("\t\tCharges: %s", acc.Charges)
			}
		}
	}
}

// CPGetCmd gets one specific counterparty.
type CPGetCmd struct {
	JSONOption
	DefaultShowOptions
	Args struct {
		ID string `required:"true" positional-arg-name:"ID" description:"ID of a counterparty. If you specify a nickname instead, the counterparty is fetched from the cache."`
	} `positional-args:"true"`
}

// Execute the counterparty get command.
func (cmd *CPGetCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	cp, err := c.GetCounterparty(cmd.Args.ID)
	if err != nil {
		return err
	}

	if cmd.JSON {
		data, _ := json.MarshalIndent(cp, "", "\t")
		slog.Msg("%s", data)
		return nil
	}

	displayCounterparty(cp, cmd.Short, cmd.Details)
	return nil
}

// CPAddCmd adds a new counterparty.
type CPAddCmd struct {
	Revolut  CPAddRevolutCmd  `command:"revolut" alias:"rev" description:"Add an existing Revolut user as a new counterparty."`
	External CPAddExternalCmd `command:"external" alias:"ex" description:"Add an external bank account as a new counterparty."`
}

// CPAddRevolutCmd adds a Revolut counterparty.
type CPAddRevolutCmd struct {
	Business bool   `short:"b" long:"business" description:"The counterparty is a business account. Will be personal if unspecified."`
	Name     string `short:"n" long:"name" description:"Name for a personal account. Required." value-name:"<PERSONAL NAME>"`
	Phone    string `short:"p" long:"phone" description:"Phone number for a personal account. Required." value-name:"<PHONE NUMBER>"`
	Email    string `short:"e" long:"email" description:"E-mail for an admin of a business account. Required." value-name:"<E-MAIL>"`
	Args     struct {
		Nick string `required:"true" positional-arg-name:"nickname" description:"Nickname for reference in commands."`
	} `positional-args:"true"`
}

// Execute the add Revolut counterparty command.
func (cmd *CPAddRevolutCmd) Execute(args []string) error {
	cp := revolut.InternalCounterparty{}
	cp.Email = cmd.Email
	cp.Name = cmd.Name
	cp.Phone = cmd.Phone
	if cmd.Business {
		if len(cmd.Email) == 0 {
			return errors.New("e-mail is required for business accounts")
		}
		cp.ProfileType = "business"
	} else {
		if len(cmd.Name) == 0 {
			return errors.New("a name is required for a personal account")
		}
		if len(cmd.Phone) == 0 {
			return errors.New("a phone number is required for a personal account")
		}
		cp.ProfileType = "personal"
	}

	c, err := newClient()
	if err != nil {
		return err
	}

	resp, err := c.AddRevolutCounterparty(cp)
	if err != nil {
		return err
	}

	slog.Msg("Counterparty %s added successfully.", resp.ID)
	return nil
}

// CPAddExternalCmd adds an external (non-Revolut) counterparty.
type CPAddExternalCmd struct {
	Business bool `short:"b" long:"business" description:"The counterparty is a business account. Will be personal if unspecified."`
	Args     struct {
		Nick     string `required:"true" positional-arg-name:"NICKNAME" description:"Nickname for reference in commands."`
		Filename string `required:"true" positional-arg-name:"FILENAME" description:"JSON file to load details from. Use the 'json' tool command to show an example to start from."`
	} `positional-args:"true"`
}

// Execute the add External counterparty command.
func (cmd *CPAddExternalCmd) Execute(args []string) error {
	if !cross.FileExists(cmd.Args.Filename) {
		return errors.New("no such file: " + cmd.Args.Filename)
	}

	data, err := ioutil.ReadFile(cmd.Args.Filename)
	if err != nil {
		return err
	}

	cp := revolut.ExternalCounterparty{}
	err = json.Unmarshal(data, &cp)
	if err != nil {
		return err
	}

	c, err := newClient()
	if err != nil {
		return err
	}

	res, err := c.AddExternalCounterparty(cp)
	if err != nil {
		return err
	}

	slog.Msg("Counterparty %s added successfully.", res.ID)
	return nil
}

// CPDeleteCmd deletes a counterparty.
type CPDeleteCmd struct {
	Args struct {
		ID string `required:"true" positional-arg-name:"ID" description:"UUID of counterparty to delete."`
	} `positional-args:"true"`
}

// Execute the delete command.
func (cmd *CPDeleteCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		slog.Msg("ERROR: %s", err.Error())
		return err
	}

	err = c.DeleteCounterparty(cmd.Args.ID)
	if err != nil {
		return err
	}

	slog.Msg("Counterparty deleted.")
	return nil
}

// CPUpdateCmd refreshes the counterparty cache.
type CPUpdateCmd struct{}

// Execute the update command.
func (cmd *CPUpdateCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	cache := &CounterpartyCache{}
	list, err := c.GetCounterparties()
	if err != nil {
		return err
	}

	for _, cp := range list {
		cache.Set(cp.ID, cp)
	}
	cache.Save()
	return nil
}
