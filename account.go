package revolut

import (
	"encoding/json"
	"errors"
)

// Account holds one business account, or the response from adding a counterparty.
type Account struct {
	// ID is the UUID, and is always available.
	ID string `json:"id"`
	// Name is the display name. Not used in counterparty responses.
	Name string `json:"name,omitempty"`
	// Balance is not used in counterparty responses.
	Balance float64 `json:"balance,omitempty"`
	// Currency is always available.
	Currency string `json:"currency"`
	// State is not used in counterparty responses.
	State string `json:"state,omitempty"`
	// Public is not used in counterparty responses.
	Public bool `json:"public,omitempty"`
	// Created is an ISO date/time. Not used in counterparty responses.
	Created string `json:"created_at,omitempty"`
	// Updated is an ISO date/time. Mot used in counterparty responses.
	Updated string `json:"updated_at,omitempty"`
	// Type is only used in counterparty responses.
	Type string `json:"type,omitempty"`
}

// BankDetails can be retrieved for an account ID.
type BankDetails struct {
	// IBAN if applicable.
	IBAN string `json:"iban,omitempty"`
	// BIC if applicable.
	BIC string `json:"bic,omitempty"`
	// AccountNo if applicable.
	AccountNo string `json:"account_no,omitempty"`
	// SortCode if applicable.
	SortCode string `json:"sort_code,omitempty"`
	// RoutingNo for the bank.
	RoutingNo string `json:"routing_number,omitempty"`
	// Beneficiary name.
	Beneficiary string `json:"beneficiary"`
	// BeneficiaryAddress is a sub-structure.
	Address Address `json:"beneficiary_address"`
	// Country of the bank. This is a two-letter ISO code.
	Country string `json:"bank_country"`
	// Pooled or unique status.
	Pooled bool `json:"pooled"`
	// UniqueReference of the pooled account.
	UniqueReference string `json:"unique_reference,omitempty"`
	//Schemes is one of: chaps, bacs, faster_payments, sepa, swift, ach
	Schemes []string `json:"schemes"`
	// EstimatedTime for transfers.
	EstimatedTime EstimatedTime `json:"estimated_time"`
}

// EstimatedTime for transfers.
type EstimatedTime struct {
	// Unit is "days" or "hours".
	Unit string `json:"unit"`
	// Max days or hours.
	Max int `json:"max"`
	// Min days or hours.
	Min int `json:"min"`
}

// GetAccounts lists the accounts for a given API key.
func (c *Client) GetAccounts() ([]Account, error) {
	contents, code, err := c.GetJSON(epAccounts)

	if code != 200 {
		err = errors.New(codeToError(code))
		return nil, err
	}

	var data []Account
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetAccount retrieves the basic information for a given account ID.
func (c *Client) GetAccount(id string) (*Account, error) {
	var acc Account
	contents, code, err := c.GetJSON(epAccounts + "/" + id)
	c.ErrorCode = code
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, &acc)
	return &acc, err
}
