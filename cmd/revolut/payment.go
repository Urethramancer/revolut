package main

// PaymentCmd contains all the payment and transaction commands.
type PaymentCmd struct {
	List PayListCmd `command:"list" alias:"ls" description:"List payment/transaction history with optional filters."`
}

// PayListCmd shows payments and/or internal transactions.
type PayListCmd struct {
	Short bool `short:"s" description:"Shorten IDs for display purposes."`
}
