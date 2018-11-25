package main

// PaymentCmd contains all the payment and transaction commands.
type PaymentCmd struct {
}

// PayListCmd lists payments and/or internal transactions.
type PayListmd struct {
	Short bool `short:"s" description:"Shorten IDs for display purposes."`
}
