package main

import (
	"os"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

//
// Bank details cache
//

// DetailsCache holds previously seen bank details.
type DetailsCache map[string]*[]revolut.BankDetails

// HasID convenience function.
func (c *DetailsCache) HasID(id string) bool {
	if len(*c) == 0 {
		return false
	}

	_, ok := (*c)[id]
	return ok
}

// Get the details for an ID.
func (c *DetailsCache) Get(id string) *[]revolut.BankDetails {
	return (*c)[id]
}

// Set the details list for an ID.
func (c *DetailsCache) Set(id string, list *[]revolut.BankDetails) {
	(*c)[id] = list
}

// Add slice of band details to an ID.
func (c *DetailsCache) Add(id string, det *[]revolut.BankDetails) {
	(*c)[id] = det
}

// loadBankDetails loads the cached account details.
func loadBankDetails() *DetailsCache {
	fn := cross.ConfigName(BankDetailsFile)
	cache := DetailsCache{}
	err := LoadJSON(fn, cache)
	if err != nil {
		slog.Warn("Couldn't load bank details cache: %s. Proceeding with clean slate.", err.Error())
	}

	return &cache
}

// saveBankDetails saves the account details cache.
func saveBankDetails(cache *DetailsCache) {
	fn := cross.ConfigName(BankDetailsFile)
	err := SaveJSON(fn, cache)
	if err != nil {
		slog.Error("Error saving account cache: ", err.Error())
		os.Exit(2)
	}
}

//
// Counterparty cache
//

// CounterpartyCache holds previously seen counterparties.
type CounterpartyCache map[string]*[]revolut.Counterparty

// HasID convenience function.
func (c *CounterpartyCache) HasID(id string) bool {
	if len(*c) == 0 {
		return false
	}

	_, ok := (*c)[id]
	return ok
}

// Get the details for an ID.
func (c *CounterpartyCache) Get(id string) *[]revolut.Counterparty {
	return (*c)[id]
}

// Set the details list for an ID.
func (c *CounterpartyCache) Set(id string, list *[]revolut.Counterparty) {
	(*c)[id] = list
}

// Add slice of band details to an ID.
func (c *CounterpartyCache) Add(id string, cp *[]revolut.Counterparty) {
	(*c)[id] = cp
}

// loadAccounts loads the cached account details.
func loadCounterparties() *CounterpartyCache {
	fn := cross.ConfigName(CounterpartiesFile)
	cache := CounterpartyCache{}
	err := LoadJSON(fn, cache)
	if err != nil {
		slog.Warn("Couldn't load counterparty cache: %s. Proceeding with clean slate.", err.Error())
	}

	return &cache
}

// saveAccounts saves the account details cache.
func saveCounterparties(cache *CounterpartyCache) {
	fn := cross.ConfigName(CounterpartiesFile)
	err := SaveJSON(fn, cache)
	if err != nil {
		slog.Error("Error saving counterparty cache: ", err.Error())
		os.Exit(2)
	}
}
