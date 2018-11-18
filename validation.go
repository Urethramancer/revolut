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
