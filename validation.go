package revolut

// ValidKey checks that the supplied string conforms roughly to a valid Revolut API key's format.
func ValidKey(s string) bool {
	// Too short, so definitely wrong.
	if len(s) < 40 {
		return false
	}

	switch s[0:5] {
	case "prod_":
		return true
	case "sand_":
		return true
	}

	return false
}

// ValidTransactionType checks if a type filter is correct.
func ValidTransactionType(t string) bool {
	switch t {
	case "atm", "card_payment", "card_refund", "card_chargeback", "card_credit", "exchange", "transfer", "loan", "fee", "refund", "topup", "topup_return", "tax", "tax_refund":
		return true
	}

	return false
}
