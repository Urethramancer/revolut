package main

// DefaultShowOptions are used in every command which lists things.
type DefaultShowOptions struct {
	Short   bool `short:"s" long:"shorten" description:"Shorten IDs for display purposes."`
	Details bool `short:"d" long:"details" description:"Show detailed information."`
}

// CurrencyOptions are used by commands with currency filters.
type CurrencyOptions struct {
	Currencies string `short:"c" description:"List only this comma-separated list of currencies."`
}
