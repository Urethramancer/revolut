# Revolut for Business (unofficial) [![Build Status](https://travis-ci.org/Urethramancer/revolut.svg?branch=master)](https://travis-ci.org/Urethramancer/revolut) [![GoDoc](https://godoc.org/github.com/Urethramancer/revolut?status.svg)](https://godoc.org/github.com/Urethramancer/revolut) [![Go Report Card](https://goreportcard.com/badge/github.com/Urethramancer/revolut)](https://goreportcard.com/report/github.com/Urethramancer/revolut)
Go SDK for accessing the Revolut for Business sandbox and production APIs.


## What does it do?
This package gives you an SDK to access the [Revolut for Business](https://business.revolut.com/signin#login) API, as described in [their documentation](https://revolutdev.github.io/business-api/#api-v1-0-introduction). You can:
- Retrieve a list of accounts or individual account information
- Add, remove and view counterparties (businesses and individuals you send money to)
- Initiate/schedule/cancel payments and check their status
- Get transaction history
- Transfer money between accounts
- Add web-hooks and unmarshal their data


## Requirements
Go 1.11 or newer is all you need to get started.

The required packages for the command line tool should install automatically when you `go get` this:
- github.com/jessevdk/go-flags
- github.com/Urethramancer/cross
- github.com/Urethramancer/slog


## Supported operating systems
Most or all systems supported by Go should be able to use this. Testing is mainly done on Unix-like systems, and the CLI tool does not have proper Windows support yet. It will run, but probably won't store its coniguration where you'd like it.


## Installing it

```sh
go get github.com/Urethramancer/revolut
```


## Using it

### Retrieve a list of accounts

To retrieve a slice of accounts:
```go
c := revolut.NewClient("prod_wibblewobblewomble") // Replace this string with your API key
list, _ := c.GetAccounts()
```

The NewClient() function will select the correct API to use based on the format of the key. Errors will be returned if the key looks wrong or when unable to connect.


### Retrieve banking details for an account

Continuing from the code above, you could do this to get a slice of bank details for the first account:
```go
details := c.GetAccountDetails(list[0].ID)
```

### Transferring between own accounts

This transfers money within accounts on one API key:
```go
// Custom request ID, source account, destination account, currency, optional message, amount
resp, err := c.Transfer("2cbd8cd60026c6b03f8c576c81f9a39dcf251e2b.", "374e6066-3830-4000-abbf-b2e240349000", "a4f9667b-4435-4dea-bad5-74cc4de9c2bf", "GBP, "I like GBP better", 200000)
```

The response is a TransferResponse, containing a new UUID for this request and its status. The reason field will contain an explanation if status is anything but "completed". Note that the currency must match the currency of the receiving account. You can't transfer GBP to an account set to USD.

### Retrieve counterparties

Paying external recipients (counterparties) can be done like this:
```go
resp, err := c.Pay(id, cmd.Args.Account, cmd.Args.Counterparty, cmd.RecAccount, cmd.Args.Currency, cmd.Reference, cmd.ScheduleTime, cmd.Args.Amount)
```

The response is a PaymentResponse, similar to transfers.

### List transactions

All transfers and payments can be retrieved, optionally filtered by start and end dates, the receiving counterparty, a type of transaction and maximum number to show:
```go
tr, err := c.GetTransactions("fee", "2018-11-20", "2018-12-02", "", 500) // Only fee transactions between two dates, any counterparty, max 500
```

This returns a slice of TransactionStatus structures, each containing a slice of Legs with information about the journey of the transaction.
