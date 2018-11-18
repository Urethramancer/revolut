package main

import (
	"github.com/Urethramancer/cross"
	"github.com/jessevdk/go-flags"
)

var O struct {
	//
	// Tool commands
	//
	Version      VersionCmd      `command:"version" alias:"ver" description:"Show version and exit."`
	AppConfig    AppConfigCmd    `command:"config" alias:"cfg" description:"Application configuration."`
	Account      AccountCmd      `command:"account" alias:"acc" description:"Account details."`
	Counterparty CounterpartyCmd `command:"counterparty" alias:"cp" description:"Counterparty listing and management."`
	Transaction  TransactionCmd  `command:"transaction" alias:"tr" description:"Transaction listing and management."`
	Payment      PaymentCmd      `command:"payment" alias:"pay" description:"Payment creation."`
	Webhooks     WebhooksCmd     `command:"webhooks" alias:"web" description:"Webhook listing and management."`
}

// Execute creates a new configuration file.
func init() {
	LoadConfig()
}

func main() {
	cross.SetConfigPath(programName)
	flags.Parse(&O)
}
