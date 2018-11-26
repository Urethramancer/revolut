package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Urethramancer/revolut"
	"github.com/Urethramancer/slog"
)

// AccountCmd holds tool commands for account viewing and management.
type AccountCmd struct {
	List   AccListCmd   `command:"list" alias:"ls" description:"List accounts. They will be loaded from the cache if available."`
	Show   AccShowCmd   `command:"show" description:"Show one specific account by UUID. It will be loaded from the cache if available."`
	Update AccUpdateCmd `command:"update" alias:"up" description:"Refresh the bank details cache."`
}

//
// Listing accounts.
//

// AccListCmd is empty.
type AccListCmd struct {
	DefaultShowOptions
	Currencies string `short:"c" description:"Show only this comma-separated list of currencies." value-name:"<CURRENCY,...>"`
}

// Execute lists the user's accounts.
func (cmd *AccListCmd) Execute(args []string) error {
	cache := loadBankDetails()

	c, err := newClient()
	if err != nil {
		return err
	}

	accounts, err := c.GetAccounts()
	if err != nil {
		return err
	}

	if len(*cache) == 0 {

		if len(accounts) == 0 {
			slog.Msg("No accounts to list.")
			return nil
		}

		for _, acc := range accounts {
			det, err := c.GetAccountDetails(acc.ID)
			if err != nil {
				return err
			}
			cache.Set(acc.ID, det)
		}
		saveBankDetails(cache)
	}

	var list []string
	for k := range *cache {
		list = append(list, k)
	}
	sort.Strings(list)

	slog.Msg("Accounts:")
	for _, acc := range accounts {
		id := acc.ID
		if cmd.Short {
			a := strings.Split(id, "-")
			id = a[4]
		}
		name := acc.Name
		if len(name) == 0 {
			name = "<unnamed>"
		}

		if !shouldDisplayCurrency(acc.Currency, cmd.Currencies) {
			continue
		}

		slog.Msg("%s (%s): %s - %f %s", id, acc.State, name, acc.Balance, acc.Currency)
		if cmd.Details {
			showDetails(cache.Get(acc.ID))
		}
	}
	return nil
}

func showDetails(det *[]revolut.BankDetails) {
	for _, d := range *det {
		prDet("Account number:", d.AccountNo)
		prDet("Sort code:", d.SortCode)
		prDet("IBAN:", d.IBAN)
		prDet("BIC:", d.BIC)
		prDet("Beneficiary:", d.Beneficiary)
		prDet("Beneficiary address:\n\t\t", d.Address.Street1)
		prDet("\t", d.Address.Street2)
		prDet("\t", d.Address.Postcode)
		prDet("\t", d.Address.City)
		prDet("\t", d.Address.Region)
		prDet("\t", d.Address.Country)
		prDet("BIC:", d.BIC)
		prDet("Bank country:", d.Country)
		s := strings.Join(d.Schemes, ", ")
		prDet("Schemas:", s)
		p := fmt.Sprintf("%t", d.Pooled)
		prDet("Pooled:", p)
		t := fmt.Sprintf("%d-%d %s", d.EstimatedTime.Min, d.EstimatedTime.Max, d.EstimatedTime.Unit)
		prDet("Estimated time:", t)
		slog.Msg("")
	}
}

func prDet(pre, x string) {
	if len(x) > 0 {
		slog.Msg("\t%s %s", pre, x)
	}
}

// AccShowCmd shows one account.
type AccShowCmd struct {
	DefaultShowOptions
	CurrencyOptions
	Args struct {
		ID string `required:"true" positional-arg-name:"ID" description:"UUID of account to show."`
	} `positional-args:"true"`
}

// Execute the single-account listing.
func (cmd *AccShowCmd) Execute(args []string) error {
	details := loadBankDetails()
	var det *[]revolut.BankDetails
	var err error
	if details.HasID(cmd.Args.ID) {
		det = details.Get(cmd.Args.ID)
	} else {
		var c *revolut.Client
		c, err = newClient()
		if err != nil {
			return err
		}

		det, err = c.GetAccountDetails(cmd.Args.ID)
		if err != nil {
			return err
		}

		// Save it to the cache
		details.Set(cmd.Args.ID, det)
		saveBankDetails(details)
	}

	showDetails(det)
	return nil
}

//
// Account caching.
//

// AccUpdateCmd fetches a list of accounts and stores the bank details for each locally for faster lookup.
type AccUpdateCmd struct{}

// Execute the account details update.
func (cmd *AccUpdateCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	list, err := c.GetAccounts()
	if err != nil {
		return err
	}

	cache := &DetailsCache{}
	for _, acc := range list {
		if !cache.HasID(acc.ID) {
			det, err := c.GetAccountDetails(acc.ID)
			if err != nil {
				return err
			}
			cache.Set(acc.ID, det)
		}
	}

	saveBankDetails(cache)
	return nil
}
