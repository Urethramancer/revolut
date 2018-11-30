package main

import (
	"errors"
	"sort"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/revolut"
)

//
// Account cache
//

// AccountCache stores a list of accounts.
type AccountCache map[string]revolut.Account

// IsEmpty checks if the map is empty.
func (c *AccountCache) IsEmpty() bool {
	return len(*c) == 0
}

// Contains checks if the cache contains the given ID.
func (c *AccountCache) Contains(id string) bool {
	_, ok := (*c)[id]
	return ok
}

// Get returns an account by ID.
func (c *AccountCache) Get(id string) interface{} {
	return (*c)[id]
}

// Set stores an account by ID.
func (c *AccountCache) Set(id string, acc revolut.Account) {
	(*c)[id] = acc
}

// SortedList of accounts.
func (c *AccountCache) SortedList() []string {
	var l []string
	for k := range *c {
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
type DetailsCache map[string][]revolut.BankDetails

// HasID convenience function.
func (c *DetailsCache) HasID(id string) bool {
	_, ok := (*c)[id]
	return ok
}

// Get the details for an ID.
func (c *DetailsCache) Get(id string) []revolut.BankDetails {
	return (*c)[id]
}

// Set the details list for an ID.
func (c *DetailsCache) Set(id string, list []revolut.BankDetails) {
	(*c)[id] = list
}

// Load from file.
func (c DetailsCache) Load() error {
	fn := cross.ConfigName(DetailsFile)
	if !cross.FileExists(fn) {
		return errors.New("no existing bank details cache file")
	}

	return LoadJSON(fn, &c)
}

// Save to file.
func (c DetailsCache) Save() error {
	return SaveJSON(cross.ConfigName(DetailsFile), c)
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

// Load from file.
func (c CounterpartyCache) Load() error {
	fn := cross.ConfigName(CounterpartiesFile)
	if !cross.FileExists(fn) {
		return errors.New("no existing counterparty cache file")
	}

	return LoadJSON(fn, &c)
}

// Save to file.
func (c CounterpartyCache) Save() error {
	return SaveJSON(cross.ConfigName(CounterpartiesFile), c)
}
