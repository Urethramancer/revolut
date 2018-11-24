# Revolut for Business (unofficial) [![Build Status](https://travis-ci.org/Urethramancer/revolut.svg?branch=master)](https://travis-ci.org/Urethramancer/revolut) [![GoDoc](https://godoc.org/github.com/Urethramancer/revolut?status.svg)](https://godoc.org/github.com/Urethramancer/revolut)
Go SDK for accessing the Revolut for Business sandbox and production APIs.


## What does it do?
This package gives you an SDK to access the [Revolut for Business](https://business.revolut.com/signin#login) API, as described in [their documentation](https://revolutdev.github.io/business-api/#api-v1-0-introduction). You can:
- Retrieve a list of accounts or individual account information
- Add, remove and view counterparties (businesses and individuals you send money to)
- Initiate/schedule/cancel payments and check their status
- Get a list of transactions
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
