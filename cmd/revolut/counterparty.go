package main

import (
	"errors"
	"strings"
	"time"

	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

type CounterpartyCmd struct {
	List   CPListCmd   `command:"list" alias:"ls" description:"List counterparties."`
	Add    CPAddCmd    `command:"add" description:"Add a new counterparty."`
	Delete CPDeleteCmd `command:"delete" alias:"del" alias:"rm" description:"Delete a counterparty."`
}

// CPListCmd shows a list of counterparties.
type CPListCmd struct {
	Short   bool `short:"s" description:"Shorten IDs for display purposes."`
	Details bool `short:"d" description:"Show details for each account."`
}

// Execute the counterparty list command.
func (cmd *CPListCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		slog.Msg("ERROR: %s", err.Error())
		return err
	}

	list, err := c.GetCounterparties()
	if err != nil {
		return err
	}

	if len(list) == 0 {
		slog.Msg("No counterparties to list.")
		return nil
	}

	slog.Msg("Counterparties:")
	for _, cp := range list {
		id := cp.ID
		if cmd.Short {
			a := strings.Split(id, "-")
			id = a[len(a)-1]
		}
		switch cp.Type {
		case "personal":
			slog.Msg("%s (%s): %s (%s), Phone: %s, updated %s", id, cp.Type, cp.Name, cp.Country, cp.Phone, cp.UpdatedAt.Format(time.RFC822))
		case "business":
			slog.Msg("%s (%s): (%s), updated %s", id, cp.Type, cp.Country, cp.UpdatedAt.Format(time.RFC822))
		}
		if !cmd.Details {
			continue
		}
		slog.Msg("\tBank details:")
		for _, acc := range cp.Accounts {
			id = acc.ID
			if cmd.Short {
				a := strings.Split(id, "-")
				id = a[len(a)-1]
			}
			slog.Msg("\t%s (%s, %s)", id, acc.Type, acc.Currency)
		}
	}
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
	cp := revolut.RevolutCounterparty{}
	if cmd.Business {
		if len(cmd.Email) == 0 {
			return errors.New("e-mail is required for business accounts")
		}
		cp.ProfileType = "business"
		cp.Email = cmd.Email
	} else {
		if len(cmd.Name) == 0 {
			return errors.New("a name is required for a personal account")
		}
		if len(cmd.Phone) == 0 {
			return errors.New("a phone number is required for a personal account")
		}
		cp.ProfileType = "personal"
		cp.Name = cmd.Name
		cp.Phone = cmd.Phone
	}

	c, err := newClient()
	if err != nil {
		return err
	}

	resp, err := c.AddRevolutCounterparty(cp)
	if err != nil {
		return err
	}

	slog.Msg("%#v", resp)
	slog.Msg("Counterparty added successfully.")
	return nil
}

// CPAddExternalCmd adds an external (non-Revolut) counterparty.
type CPAddExternalCmd struct{}

// Execute the add External counterparty command.
func (cmd *CPAddExternalCmd) Execute(args []string) error {
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
