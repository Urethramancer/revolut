package main

import (
	"os"
	"sort"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

//
// Account cache
//

// AccountCache stores a list of accounts.
type AccountCache map[string]revolut.Account

// IsEmpty checks if the map is empty.
func (c AccountCache) IsEmpty() bool {
	return len(c) == 0
}

// Contains checks if the cache contains the given ID.
func (c AccountCache) Contains(id string) bool {
	_, ok := c[id]
	return ok
}

// Get returns an account by ID.
func (c AccountCache) Get(id string) interface{} {
	return c[id]
}

// Set stores an account by ID.
func (c AccountCache) Set(id string, acc revolut.Account) {
	c[id] = acc
}

// SortedList of accounts.
func (c AccountCache) SortedList() []string {
	var l []string
	for k := range c {
		l = append(l, k)
	}
	sort.Strings(l)
	return l
}

// Load from file.
func (c AccountCache) Load(path string) error {
	return LoadJSON(path, &c)
}

// Save to file.
func (c AccountCache) Save(path string) error {
	return SaveJSON(path, c)
}

//
// Bank details cache
//

// DetailsCache holds previously seen bank details.
type DetailsCache map[string]*[]revolut.BankDetails

// IsEmpty checks if the map is empty.
func (c *DetailsCache) IsEmpty() bool {
	return len(*c) == 0
}

// HasID convenience function.
func (c *DetailsCache) HasID(id string) bool {
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

// Load from file.
func (c DetailsCache) Load(path string) error {
	return LoadJSON(path, &c)
}

// Save to file.
func (c DetailsCache) Save(path string) error {
	return SaveJSON(path, c)
}

//
// Counterparty cache
//

// CounterpartyCache holds previously seen counterparties.
type CounterpartyCache map[string]*revolut.Counterparty

// HasID convenience function.
func (c *CounterpartyCache) HasID(id string) bool {
	if len(*c) == 0 {
		return false
	}

	_, ok := (*c)[id]
	return ok
}

// Get the cached counterparty for an ID.
func (c *CounterpartyCache) Get(id string) *revolut.Counterparty {
	return (*c)[id]
}

// Set the cached counterparty for an ID.
func (c *CounterpartyCache) Set(id string, cp *revolut.Counterparty) {
	(*c)[id] = cp
}

// loadAccounts loads the cached account details.
func loadCounterparties() *CounterpartyCache {
	fn := cross.ConfigName(CounterpartiesFile)
	cache := &CounterpartyCache{}
	err := LoadJSON(fn, cache)
	if err != nil {
		slog.Warn("Couldn't load counterparty cache: %s. Proceeding with clean slate.", err.Error())
	}

	return cache
}

// saveAccounts saves the account details cache.
func saveCounterparties(cache *CounterpartyCache) {
	fn := cross.ConfigName(CounterpartiesFile)
	err := SaveJSON(fn, &cache)
	if err != nil {
		slog.Error("Error saving counterparty cache: ", err.Error())
		os.Exit(2)
	}
}
