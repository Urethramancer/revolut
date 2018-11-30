package main

import (
	"fmt"
	"strings"

	"github.com/Urethramancer/cross"

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
	var err error
	acache := &AccountCache{}
	acachename := cross.ConfigName(AccountsFile)
	if cross.FileExists(acachename) {
		err = acache.Load(acachename)
		if err != nil {
			return err
		}
	}

	dcache := &DetailsCache{}
	if cross.FileExists(cross.ConfigName(DetailsFile)) {
		err = dcache.Load()
		if err != nil {
			return err
		}
	}

	if acache.IsEmpty() {
		acache, dcache, err = updateDetailsCache()
		if err != nil {
			return err
		}
	}

	if acache.IsEmpty() {
		slog.Msg("No accounts to list.")
		return nil
	}

	slog.Msg("Accounts:")
	for _, acc := range *acache {
		if len(acc.Name) == 0 {
			acc.Name = "<unnamed>"
		}

		if !shouldDisplayCurrency(acc.Currency, cmd.Currencies) {
			continue
		}

		showAccount(&acc, cmd.Short)
		if cmd.Details {
			showDetails(dcache.Get(acc.ID))
		}
	}
	return nil
}

func showAccount(acc *revolut.Account, short bool) {
	if short {
		acc.ID = shortUUID(acc.ID)
	}
	slog.Msg("%s (%s): %s - %f %s", acc.ID, acc.State, acc.Name, acc.Balance, acc.Currency)
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
	CurrenciesOption
	Args struct {
		ID string `required:"true" positional-arg-name:"ID" description:"UUID of account to show."`
	} `positional-args:"true"`
}

// Execute the single-account listing.
func (cmd *AccShowCmd) Execute(args []string) error {
	dcache := DetailsCache{}
	err := dcache.Load()
	if err != nil {
		slog.Warn("Warning: %s", err.Error())
	}

	var det *[]revolut.BankDetails
	if dcache.HasID(cmd.Args.ID) {
		det = dcache.Get(cmd.Args.ID)
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
		dcache.Set(cmd.Args.ID, det)
		err = dcache.Save()
		if err != nil {
			slog.Warn("Warning: %s", err.Error())
		}
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
	_, _, err := updateDetailsCache()
	return err
}

func updateDetailsCache() (*AccountCache, *DetailsCache, error) {
	var err error
	acache := &AccountCache{}
	dcache := &DetailsCache{}

	c, err := newClient()
	if err != nil {
		return acache, dcache, err
	}

	accounts, err := c.GetAccounts()
	if err != nil {
		return acache, dcache, err
	}

	for _, acc := range accounts {
		acache.Set(acc.ID, acc)
	}
	err = acache.Save(cross.ConfigName(AccountsFile))
	if err != nil {
		return acache, dcache, err
	}

	for _, acc := range accounts {
		if !dcache.HasID(acc.ID) {
			var det *[]revolut.BankDetails
			det, err = c.GetAccountDetails(acc.ID)
			if err != nil {
				return acache, dcache, err
			}
			dcache.Set(acc.ID, det)
		}
	}

	err = dcache.Save()
	return acache, dcache, err
}
