package main

import "github.com/Urethramancer/revolut"

type DetailsMap map[string]*[]revolut.BankDetails

// HasID convenience function.
func (d *DetailsMap) HasID(id string) bool {
	if len(*d) == 0 {
		return false
	}

	_, ok := (*d)[id]
	return ok
}

// Get the details for an ID.
func (d *DetailsMap) Get(id string) *[]revolut.BankDetails {
	return (*d)[id]
}

// Add slice of band details to an ID.
func (d *DetailsMap) Add(id string, det *[]revolut.BankDetails) {
	(*d)[id] = det
}
