package revolut

import "time"

// Merchant info.
type Merchant struct {
	// Name of the merchant.
	Name string `json:"name"`
	// City of the merchant.
	City string `json:"city"`
	// Category code for the transaction.
	Category string `json:"category_code"`
	// Country is the 3-letter ISO bank country code.
	Country string `json:"country"`
}

// Leg of the transaction process.
type Leg struct {
	// ID of this leg.
	ID string `json:"leg_id"`
	// Amount of the transactiob.
	Amount float64 `json:"amount"`
	// Currency is the 3-letter ISO code for the transaction currency.
	Currency string `json:"currency"`
	// BillAmount is the amount for cross-currency transactions.
	BillAmount float64 `json:"bill_amount,omitempty"`
	// BillCurrency is the billing currency for cross-currency transactions.
	BillCurrency string `json:"bill_currency,omitempty"`
	// AccountID of the account this transaction is associated with.
	AccountID string `json:"account_id"`
	// Counterparty for this leg.
	Counterparty LegCounterparty `json:"counterparty"`
	// Description contains the leg purpose.
	Description string `json:"description"`
	// Card information only for card payments.
	Card Card `json:"card,omitempty"`
}

// LegCounterparty is a quick summary of the counterparty involved in this leg of the transaction.
type LegCounterparty struct {
	ID string `json:"id"`
	// Type is "self", "revolut" or "external".
	Type      string `json:"type"`
	AccountID string `json:"account_id"`
}

// Card is used for card payments.
type Card struct {
	// Number is the masked card number.
	Number string `json:"card_number"`
	// First name of the card holder.
	First string `json:"first_name"`
	// Last name of the card holder.
	Last string `json:"last_name"`
	// Phone number of the card holder.
	Phone string `json:"phone"`
}

// GetTransactions gets all transactions, limited by optional counts, counterparty and type.
// Returns 100 records if count is left at zero.
func (c *Client) GetTransactions(counterparty, ttype string, count int) {

}

// GetTransactionsRange gets a list of transactions for a timespan, up to an optional limit.
func (c *Client) GetTransactionsRange(from, to time.Time, counterparty, ttype string, count int) {

}

// GetTransaction returns a single transaction.
func (c *Client) GetTransaction(id string) {

}
