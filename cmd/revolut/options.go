package main

// ShortOption is the flag for brief UUIDs.
type ShortOption struct {
	Short bool `short:"s" long:"shorten" description:"Shorten IDs for display purposes."`
}

// DetailsOption is the extended information flag.
type DetailsOption struct {
	Details bool `short:"d" long:"details" description:"Show detailed information."`
}

// DefaultShowOptions are used in many commands which list things.
type DefaultShowOptions struct {
	ShortOption
	DetailsOption
}

// CurrenciesOption is used by commands with currency filters.
type CurrenciesOption struct {
	Currencies string `short:"c" description:"List only this comma-separated list of currencies."`
}

// JSONOption is used by commands which can output formatted JSON structures.
type JSONOption struct {
	JSON bool `short:"j" long:"json" description:"Print the actual JSON structure instead of formatted information."`
}

// ReferenceOption is used on transactions from your accounts.
type ReferenceOption struct {
	Reference string `short:"r" long:"reference" descripttion:"Optional reference to show on the transaction." value-name:"TEXT"`
}
