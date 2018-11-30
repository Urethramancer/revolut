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
