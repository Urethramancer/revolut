package revolut

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
