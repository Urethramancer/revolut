package revolut

import (
	"encoding/json"
	"strings"
)

// PaymentRequest for payments to counterparties.
type PaymentRequest struct {
	// RequestID for the payment, provided by the client.
	RequestID string `json:"request_id"`
	// AccountID of the account to pay from.
	AccountID string `json:"account_id"`
	// Receiver of this payment.
	Receiver Receiver `json:"receiver"`
	// Amount to pay.
	Amount float64 `json:"amount"`
	// Currency for the transaction. 3-letter ISO code.
	Currency string `json:"currency"`
	// Reference is an optional text to show on the transaction. Highly recommended.
	Reference string `json:"reference,omitempty"`
	// ScheduleTime to initiate the payment. There's no guarantee this will be fulfilled right away if the current time is used.
	ScheduleTime string `json:"schedule_for,omitempty"`
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
	// CompletedAt is the ISO time when the payment finished. Not available for asynchronous or scheduled payments.
	CompletedAt string `json:"completed_at"`
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

// Pay a Revolut account or external account.
func (c *Client) Pay(id, account, cp, cpAccount, currency, reference, schedule string, amount float64) (*PaymentResponse, error) {
	var req PaymentRequest
	req.RequestID = id
	req.AccountID = account
	req.Receiver.CounterpartyID = cp
	req.Receiver.AccountID = cpAccount
	req.Amount = amount
	req.Currency = strings.ToUpper(currency)
	req.Reference = reference
	req.ScheduleTime = schedule
	contents, code, err := c.PostJSON(epPay, req)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, jsonError(contents)
	}

	var resp PaymentResponse
	err = json.Unmarshal(contents, &resp)
	return &resp, err
}

// CancelPayment if possible.
func (c *Client) CancelPayment(id string) error {
	contents, code, err := c.Delete(epTransaction + "/" + id)
	if err != nil {
		return err
	}

	if code != 204 {
		return jsonError(contents)
	}

	return nil
}
