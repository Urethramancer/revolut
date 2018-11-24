package revolut

import (
	"encoding/json"
	"errors"
	"time"
)

// Counterparty is returned from the /counterparty and /counterparties endpoints.
type Counterparty struct {
	// ID is a UUID.
	ID string `json:"id"`
	// Name of the counterparty.
	Name string `json:"name"`
	// Phone number.
	Phone string `json:"phone"`
	// Type is "personal" or "business".
	Type string `json:"profile_type"`
	// Country is a two-letter ISO code.
	Country string `json:"country"`
	// State of the counterparty is a status string.
	State string `json:"state"`
	// CreatedAt is a timestamp for when this was added.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is a timestamp for the last change to the counterparty,
	UpdatedAt time.Time `json:"updated_at"`
	// Accounts is a list of public accounts for this counterparty.
	Accounts []CounterpartyAccount `json:"accounts"`
}

// CounterpartyAccount is embedded in Counterparty structures.
type CounterpartyAccount struct {
	// ID is the UUID.
	ID string `json:"id"`
	// Currency is a three-letter shortname.
	Currency string `json:"currency"`
	// Type of account is either "revolut" or "external".
	Type string `json:"type"`
	// Account number.
	Account string `json:"account_no,omitempty"`
	// SortCode if used.
	SortCode string `json:"sort_code,omitempty"`
	// Email for the recipient.
	Email string `json:"email,omitempty"`
	// Name of the business or person this account belongs to.
	Name string `json:"name,omitempty"`
	// Country is a two-letter ISO code.
	Country string `json:"bank_country,omitempty"`
	// Charges may be added.
	Charges bool `json:"recipient_charges,omitempty"`
}

// RevolutCounterparty is used when adding an existing Revolut account as a counterparty (i.e. contact).
type RevolutCounterparty struct {
	// ProfileType is "business" or "personal".
	ProfileType string `json:"profile_type"`
	// Name of the counterparty.
	Name string `json:"name"`
	// Phone is used with personal accounts.
	Phone string `json:"phone"`
	// Email is an optional field with the address of the admin for a business account.
	Email string `json:"email,omitempty"`
}

// CounterpartyDetails is returned after adding/removing a Revolut counterparty.
type CounterpartyResponse struct {
	// ID is the UUID of the counterparty.
	ID string `json:"id"`
	// Name of the counterparty.
	Name string `json:"name"`
	// Phone number of a personal account.
	Phone string `json:"phone"`
	// ProfileType is "business" or "personal".
	ProfileType string `json:"profile_type"`
	// Country is a 2-letter code.
	Country string `json:"bank_country"`
	// State is either "created" or "deleted".
	State string `json:"state"`
	// CreatedAt is the ISO time when the counterparty was created.
	CreatedAt string `json:"created_at"`
	// UpdateAt is the ISO time when the counterparty was last updated.
	UpdatedAt string `json:"updated_at"`
	// Accounts is a list of all the counterparty's accounts.
	Accounts []Account `json:"accounts"`
}

// ExternalCounterparty is used when adding a counterparty with a non-Revolut account.
type ExternalCounterParty struct {
	// Company must exist if Individual isn't present.
	Company string `json:"company_name,omitempty"`
	// Individual must exist if Company isn't present.
	Individual IndividualName `json:"individual_name,omitempty"`
	// BankCountry is a two-letter ISO code.
	BankCountry string `json:"bank_country"`
	// Currency is a 3-letter ISO code.
	Currency string  `json:"currency"`
	Email    string  `json:"email,omitempty"`
	Phone    string  `json:"phone,omitempty"`
	Address  Address `json:"address,omitempty"`
	// AccountNo is required for UK GBP, US USD and SWIFT accounts.
	AccountNo string `json:"account_no,omitempty"`
	// SortCode is required for UK GBP accounts.
	SortCode string `json:"sort_code,omitempty"`
	// RoutingNo is required for US USD accounts.
	RoutingNo string `json:"routing_number,omitempty"`
	// IBAN is required for IBAN countries.
	IBAN string `json:"iban,omitempty"`
	// BIC is required for IBAN/SWIFT accounts.
	BIC string `json:"bic,omitempty"`
}

// IndividualName of an account holder.
type IndividualName struct {
	// First name.
	First string `json:"first_name"`
	// Last name.
	Last string `json:"last_name"`
}

// ExternalCounterpartyResponse is returned after adding/removing an external counterparty.
type ExternalCounterpartyResponse struct {
	// ID is the UUID of the counterparty.
	ID string `json:"id"`
	// Name
	Name string `json:"name"`
	// State is either "created" or "deleted".
	State string `json:"state"`
	// CreatedAt is the ISO time/date this counterparty was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the ISO time/date this counterparty was last modified.
	UpdatedAt string `json:"updated_at"`
	// Accounts known for this counterparty.
	Accounts []ExternalAccount `json:"accounts"`
}

// ExternalAccount for counterparties in other banks.
type ExternalAccount struct {
	// ID of a counterparty's account.
	ID string `json:"id"`
	// Currency is a 3-letter ISO code.
	Currency string `json:"currency"`
	// Type is "revolut" or "external".
	Type string `json:"type"`
	// AccountNo for UK GBP, US USD and SWIFT accounts
	AccountNo string `json:"account_no,omitempty"`
	// IBAN of a foreign account.
	IBAN string `json:"iban"`
	// SortCode of a UK GBP account.
	SortCode string `json:"sort_code,omitempty"`
	// RoutingNo of a US USD account
	RoutingNo string `json:"routing_number,omitempty"`
	// BIC of an IBAN/SWIFT account.
	BIC string `json:"bic,omitempty"`
	// RecipientCharges is "no", "expected" or possible "free". TODO: Clarify with devs.
	RecipientCharges string `json:"recipient_charges"`
}

// GetCounterparties returns a list of all counterparties for an API key.
func (c *Client) GetCounterparties() ([]Counterparty, error) {
	contents, code, err := c.GetJSON(epCounterparties)

	if code != 200 {
		err = errors.New(codeToError(code))
		return nil, err
	}

	var data []Counterparty
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetCounterparty
func (c *Client) GetCounterparty(id string) {

}

// AddRevolutCounterparty adds a Revolut personal or business account as a counterparty.
func (c *Client) AddRevolutCounterparty(cp RevolutCounterparty) (*CounterpartyResponse, error) {
	contents, code, err := c.PostJSON(epCounterparty, cp)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		err = errors.New(codeToError(code))
		return nil, err
	}

	var res CounterpartyResponse
	err = json.Unmarshal(contents, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// AddExternalCounterparty adds a non-Revolut account as a counterparty.
func (c *Client) AddExternalCounterparty() {

}

// DeleteCounterparty removes a counterparty by UUID.
func (c *Client) DeleteCounterparty(id string) error {
	code, err := c.Delete(epCounterparty + "/" + id)
	if err != nil {
		return err
	}

	if code != 204 {
		err = errors.New(codeToError(code))
		return err
	}

	return nil
}
