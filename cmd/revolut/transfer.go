package main

import (
	"strings"

	"github.com/Urethramancer/slog"
)

// TransferCmd transfers money between your own Revolut for Business accounts.
type TransferCmd struct {
	ReferenceOption
	Args struct {
		From     string  `required:"true" positional-arg-name:"SOURCE ID" description:"UUID of account to transfer from."`
		To       string  `required:"true" positional-arg-name:"DEST ID" description:"UUID of account to transfer to."`
		Amount   float64 `required:"true" positional-arg-name:"AMOUNT" description:"Amount to transfer."`
		Currency string  `required:"true" positional-arg-name:"CURRENCY" description:"Currency to transfer in."`
	} `positional-args:"true"`
}

// Execute the transfer.
func (cmd *TransferCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	id := generateRequestID()
	slog.Msg("Transferring %.2f %s with ID %s.", cmd.Args.Amount, strings.ToUpper(cmd.Args.Currency), id)
	resp, err := c.Transfer(id, cmd.Args.From, cmd.Args.To, cmd.Args.Currency, cmd.Reference, cmd.Args.Amount)
	if err != nil {
		return err
	}

	slog.Msg("Created transfer %s: %s", resp.ID, resp.State)
	return nil
}
