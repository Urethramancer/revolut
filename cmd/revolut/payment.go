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
	// From date
	From string `short:"f" long:"from" description:"From date. Use YYYY-MM-DD or RFC3339." value-name:"<ISO DATE>"`
	// To date
	To string `short:"e" long:"to" description:"To date. Use YYYY-MM-DD or RFC3339." value-name:"<ISO DATE>"`
	// Counterparty UUID
	Counterparty string `short:"c" long:"counterparty" description:"UUID of counterparty to show transfers for." value-name:"<UUID>"`
	// Max transactions to show
	Max int64 `short:"m" long:"max" description:"Maximum transactions to show." default:"100" value-name:"<NUMBER>"`
	// Type of transactions to show
	Type string `short:"t" long:"type" description:"Type of transactions to show." value-name:"<TYPE>" `
}

// Execute the transaction listing.
func (cmd *PayListCmd) Execute(args []string) error {
	if cmd.To != "" && !validTransactionType(cmd.Type) {
		slog.Msg("Type must be one of atm, card_payment, card_refund, card_chargeback, card_credit, exchange, transfer, loan, fee, refund, topup, topup_return, tax or tax_refund.")
		return nil
	}

	c, err := newClient()
	if err != nil {
		return err
	}

	tr, err := c.GetTransactions(cmd.Type, cmd.From, cmd.To, cmd.Counterparty, cmd.Max)
	if err != nil {
		return err
	}

	if len(tr) == 0 {
		slog.Msg("No transactions to show.")
		return nil
	}

	for _, t := range tr {
		displayTransaction(t, cmd.Short, cmd.Details)
	}
	return nil
}

func validTransactionType(t string) bool {
	switch t {
	case "atm", "card_payment", "card_refund", "card_chargeback", "card_credit", "exchange", "transfer", "loan", "fee", "refund", "topup", "topup_return", "tax", "tax_refund":
		return true
	}

	return false
}

func displayTransaction(t revolut.TransactionStatus, short, details bool) {
	id := t.ID
	if short {
		id = shortUUID(id)
	}

	legs := len(t.Legs)
	slog.Msg("%s (%s), %d leg(s), %s: %.2f %s, %s", id, t.Type, legs, t.State, t.Legs[0].Amount, t.Legs[0].Currency, t.Legs[0].Description)
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
