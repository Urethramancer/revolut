package revolut

// Address is an account-holder's address.
type Address struct {
	// Street1 is line 1 of the address.
	Street1 string `json:"street_line1"`
	// Street2 is line 2 of the address.
	Street2 string `json:"street_line2"`
	// Region of the beneficiary.
	Region string `json:"region"`
	// City of the beneficiary.
	City string `json:"city"`
	// Country of the beneficiary. This is a two-letter ISO code.
	Country string `json:"country"`
	// Postcode of the beneficiary.
	Postcode string `json:"postcode"`
}
