package main

import (
	"github.com/Urethramancer/cross"
	"github.com/jessevdk/go-flags"
)

// O holds option flags and arguments.
var O struct {
	//
	// Tool commands
	//
	Version      VersionCmd      `command:"version" alias:"ver" description:"Show version and exit."`
	AppConfig    AppConfigCmd    `command:"config" alias:"cfg" description:"Application configuration."`
	Account      AccountCmd      `command:"account" alias:"acc" description:"Account details."`
	Counterparty CounterpartyCmd `command:"counterparty" alias:"cp" description:"Counterparty listing and management."`
	Transfer     TransferCmd     `command:"transfer" alias:"tr" description:"Transfer between your own accounts."`
	Payment      PaymentCmd      `command:"payments" alias:"pay" description:"Payments and transactions."`
	Webhooks     WebhooksCmd     `command:"webhooks" alias:"web" description:"Webhook listing and management."`
	JSON         JSONCmd         `command:"json" description:"Print example data structures for JSON input."`
	Cache        CacheCmd        `command:"cache" description:"Cache manipulation."`
}

// Execute creates a new configuration file.
func init() {
	LoadConfig()
}

func main() {
	cross.SetConfigPath(programName)
	flags.Parse(&O)
}
