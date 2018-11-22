package revolut

import "time"

// TransferRequest for money transfers within a business.
type TransferRequest struct {
	// Amount to transfer.
	Amount float64 `json:"amount"`
	// ID of the request, provided by the client.
	ID string `json:"request_id"`
	// SourceID of the account to transfer from.
	SourceID string `json:"source_account_id"`
	// TargetID of the account to transfer to.
	TargetID string `json:"target_account_id"`
	// Currency to use.
	Currency string `json:"currency"`
	// Reference is an optional text to show on the transaction. Highly recommended.
	Reference string `json:"reference"`
}

// TransferResponse to a transfer request.
type TransferResponse struct {
	// ID of the created transaction.
	ID string `json:"id"`
	// State of the transaction. One of the following: "pending", "completed", "declined" or "failed".
	State string `json:"state"`
	// CreatedAT ISO date/time.
	CreatedAt string `json:"created_at"`
	// CompletedAt ISO date/time.
	CompletedAt string `json:"completed_at"`
}

//  RequestTransfer transfers money between two of a business' accounts in the same currency.
func RequestTransfer() {

}

// PaymentRequest for payments to counterparties.
type PaymentRequest struct {
	// RequestID for the payment, provided by the client.
	RequestID string `json:"request_id"`
	// AccountID of the account to pay from.
	AccountID string `json:"account_id"`
	// Amount to pay.
	Amount float64 `json:"amount"`
	// Currency for the transaction. 3-letter ISO code.
	Currency string `json:"currency"`
	// Reference is an optional text to show on the transaction. Highly recommended.
	Reference string `json:"reference,omitempty"`
	// ScheduleTime to initiate the payment. There's no guarantee this will be fulfilled right away if the current time is used.
	ScheduleTime string `json:"schedule_for"`
}

// Receiver of a payment.
type Receiver struct {
	// CounterpartyID for the receiving party.
	CounterpartyID string `json:"counterparty_id"`
	// AccountID is optional.
	AccountID string `json:"account_id,omitempty"`
}

// PaymentResponse for a request.
type PaymentResponse struct {
	// ID of the created transaction.
	ID string `json:"id"`
	// State is one of "pending", "completed", "declined" or "failed".
	State string `json:"state"`
	// Reason is a code for the "declined" or "failed" states.
	Reason string `json:"reason_code"`
	// CreatedAt is the ISO time when the payment was requested.
	CreatedAt string `json:"created_at"`
	// CompletedAt is the ISO time when the payment finished.
	CompletedAt string `json:"completed_at"`
}

// Pay starts a transfer to another counterparty's account.
func Pay() {

}

// CancelTransaction tries to stop a payment in progress.
func CancelTransaction(id string) {

}

// Transaction record.
type Transaction struct {
	// ID of the transaction.
	ID string `json:"id"`
	// Type of transaction.
	Type string `json:"type"`
	// RequestID provided by the client.
	RequestID string `json:"request_id"`
	// State is one of "pending", "completed", "declined" or "failed".
	State string `json:"state"`
	// Reason code for the "declined" and "failed" states.
	Reason string `json:"reason_code"`
	// CreatedAt is an ISO date/time.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is an ISO date/time. Available when looking up transactions.
	UpdatedAt string `json:"updated_at,omitempty"`
	// CompletedAt is an ISO date/time.
	CompletedAt string `json:"completed_at"`
	// Scheduled time is an ISO date/time the transaction was scheduled to run.
	ScheduledTime string `json:"scheduled_for"`
	// Merchant info.
	Merchant Merchant `json:"merchant"`
	// Reference for the payment provided by the user.
	Reference string `json:"reference"`
	// Legs of the transaction. There will be 2 legs between your Revolut accounts and 1 in other cases.
	Legs []Leg `json:"legs"`
}

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

func (c *Client) GetTransaction(id string) {

}
