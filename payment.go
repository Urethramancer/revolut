package revolut

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
	// CompletedAt is the ISO time when the payment finished. Not available for asynchronous or scheduled payments.
	CompletedAt string `json:"completed_at"`
}

// PaymentStatus for transaction ID or request ID.
type PaymentStatus struct {
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
