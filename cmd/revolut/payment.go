package main

import (
	"fmt"
	"strings"

	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

// PaymentCmd contains all the payment and transaction commands.
type PaymentCmd struct {
	// List transfers
	List PayListCmd `command:"list" alias:"ls" description:"List payment/transaction history with optional filters."`
	// Send money
	Send PaySendCmd `command:"send" description:"Send money/pay a counterparty."`
	// Show status
	Show PayShowCmd `command:"show" alias:"status" description:"Show the status of a payment."`
	// Cancel a transaction
	Cancel PayCancelCmd `command:"cancel" description:"Cancel a scheduled payment, if possible."`
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
	if cmd.To != "" && !revolut.ValidTransactionType(cmd.Type) {
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

// PaySendCmd sends money to counterparties.
type PaySendCmd struct {
	ReferenceOption
	RecAccount   string `short:"a" long:"account" description:"Counterparty account, if necessary. This isn't required for Revolut counterparties." value-name:"ACCOUNT"`
	ScheduleTime string `short:"s" long:"schedule" description:"Scheduled time to start the payment. Use YYYY-MM-DD or ISO3339." value-name:"TIME"`
	Args         struct {
		Account      string  `required:"true" positional-arg-name:"ACCOUNT" description:"UUID of the  account to pay from."`
		Counterparty string  `required:"true" positional-arg-name:"COUNTERPARTY" description:"UUID of the receiving counterparty."`
		Amount       float64 `required:"true" positional-arg-name:"AMOUNT" description:"Amount to transfer."`
		Currency     string  `required:"true" positional-arg-name:"CURRENCY" description:"Currency to transfer in."`
	} `positional-args:"true"`
}

// Execute the payment
func (cmd *PaySendCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	id := generateRequestID()
	slog.Msg("Paying %.2f %s with ID %s.", cmd.Args.Amount, strings.ToUpper(cmd.Args.Currency), id)
	resp, err := c.Pay(id, cmd.Args.Account, cmd.Args.Counterparty, cmd.RecAccount, cmd.Args.Currency, cmd.Reference, cmd.ScheduleTime, cmd.Args.Amount)
	if err != nil {
		return err
	}

	slog.Msg("Created payment %s: %s", resp.ID, resp.State)
	return nil
}

// PayShowCmd shows a single transaction.
type PayShowCmd struct {
	ShortOption
	DetailsOption
	Args struct {
		ID string `required:"true" positional-arg-name:"TRANSACTION" description:"UUID of a transaction to view."`
	} `positional-args:"true"`
}

// Execute the transaction display.
func (cmd *PayShowCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	resp, err := c.TransactionStatus(cmd.Args.ID)
	if err != nil {
		return err
	}

	displayTransaction(*resp, cmd.Short, cmd.Details)
	return nil
}

// PayCancelCmd tries to cancel a scheduled payment.
type PayCancelCmd struct {
	Args struct {
		ID string `required:"true" positional-arg-name:"TRANSACTION" description:"UUID of a transaction to cancel."`
	} `positional-args:"true"`
}

// Execute the cancellation.
func (cmd *PayCancelCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	return c.CancelPayment(cmd.Args.ID)
}
