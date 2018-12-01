package revolut

import (
	"encoding/json"
	"strconv"
	"strings"
)

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

// TransactionStatus is returned by GetTransaction() and GetTransactions().
type TransactionStatus struct {
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
	CompletedAt string `json:"completed_at,omitempty"`
	// Scheduled time is an ISO date/time the transaction was scheduled to run.
	ScheduledTime string `json:"scheduled_for"`
	// Merchant info.
	Merchant Merchant `json:"merchant"`
	// Reference for the payment provided by the user.
	Reference string `json:"reference"`
	// Legs of the transaction. There will be 2 legs between your Revolut accounts and 1 in other cases.
	Legs []Leg `json:"legs"`
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

// TransactionCreatedEvent is posted to webooks when a new transaction has been created.
type TransactionCreatedEvent struct {
	// Event is the event name ("TransactionCreated").
	Event string `json:"event"`
	// Timestamp is the RFC3339 date and time of this event.
	Timestamp string `json:"timestamp"`
	// Data is the real payload.
	Data TransactionCreatedData `json:"data"`
}

// TransactionCreatedData is the create event payload.
type TransactionCreatedData struct {
	// ID of the transaction.
	ID string `json:"id"`
	// Type of transaction.
	Type string `json:"type"`
	// RequestID provided by the client.
	RequestID string `json:"request_id"`
	// State of the transaction.
	State string `json:"state"`
	// Reason for failure.
	Reason string `json:"reason_code,omitempty"`
	// CreatedAt timestamp.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
	// CompletedAt timestamp.
	CompletedAt string `json:"completed_at,omitempty"`
	// Reference
	Reference string `json:"reference,omitempty"`
	// Legs of the route
	Legs []Leg `json:"legs,omitempty"`
}

// TransactionChangedEvent is posted to webhooks when a transaction has updated.
type TransactionChangedEvent struct {
	// Event is the event name ("TransactionCreated").
	Event string `json:"event"`
	// Timestamp is the RFC3339 date and time of this event.
	Timestamp string `json:"timestamp"`
	// Data is the real payload.
	Data TransactionChangedData `json:"data"`
}

// TransactionChangedData is the change event payload.
type TransactionChangedData struct {
	// ID of the transaction.
	ID string `json:"id"`
	// OldState before this event.
	OldState string `json:"old_state"`
	// NewState after this event. Expected states are "pending", "completed", "declined" or "failed".
	NewState string `json:"new_state"`
}

// GetTransactions by optional filters. The from and to datescan be in the formats
// YYYY-MM-DD or RFC3339.
func (c *Client) GetTransactions(ttype, from, to, counterparty string, count int64) ([]TransactionStatus, error) {
	var args strings.Builder
	if count > 0 {
		args.WriteString("count=")
		arg := strconv.FormatInt(count, 10)
		args.WriteString(arg)
	}

	if len(ttype) > 0 {
		if args.Len() > 0 {
			args.WriteString("&")
		}
		args.WriteString("type=")
		args.WriteString(ttype)
	}

	if len(from) > 0 {
		if args.Len() > 0 {
			args.WriteString("&")
		}
		args.WriteString("from=")
		args.WriteString(from)
	}

	if len(to) > 0 {
		if args.Len() > 0 {
			args.WriteString("&")
		}
		args.WriteString("to=")
		args.WriteString(to)
	}

	var url strings.Builder
	url.WriteString(epTransactions)
	if args.Len() > 0 {
		url.WriteString("?")
		url.WriteString(args.String())
	}

	contents, code, err := c.GetJSON(url.String())
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, jsonError(contents)
	}

	var list []TransactionStatus
	err = json.Unmarshal(contents, &list)
	return list, err
}

// TransactionStatus of transfers or payments.
func (c *Client) TransactionStatus(id string) (*TransactionStatus, error) {
	contents, code, err := c.GetJSON(epTransaction + "/" + id)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, jsonError(contents)
	}

	var resp TransactionStatus
	err = json.Unmarshal(contents, &resp)
	return &resp, err
}
