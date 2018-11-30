package main

import (
	"fmt"

	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

// PaymentCmd contains all the payment and transaction commands.
type PaymentCmd struct {
	List PayListCmd `command:"list" alias:"ls" description:"List payment/transaction history with optional filters."`
}

// PayListCmd shows payments and/or internal transactions.
type PayListCmd struct {
	ShortOption
	DetailsOption
}

// Execute the transaction listing.
func (cmd *PayListCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	tr, err := c.GetTransactions("", "", "", "", 0)
	if err != nil {
		return err
	}

	for _, t := range tr {
		displayTransaction(t, cmd.Short, cmd.Details)
	}
	return nil
}

func displayTransaction(t revolut.TransactionStatus, short, details bool) {
	id := t.ID
	if short {
		id = shortUUID(id)
	}

	legs := len(t.Legs)
	slog.Msg("%s, %d leg(s), %s: %.2f %s, %s", id, legs, t.State, t.Legs[0].Amount, t.Legs[0].Currency, t.Legs[0].Description)
	if details {
		for _, l := range t.Legs {
			lid := l.ID
			if short {
				lid = shortUUID(lid)
			}
			alt := ""
			if l.BillAmount != 0 {
				alt = fmt.Sprintf(" (%.2f %s)", l.BillAmount, l.BillCurrency)
			}
			slog.Msg("\t%s: %.2f %s%s, %s", lid, l.Amount, l.Currency, alt, l.Description)
		}
	}
}
