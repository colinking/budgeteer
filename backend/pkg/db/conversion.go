package db

import (
	"github.com/colinking/budgeteer/backend/pkg/gen/userpb"
	"github.com/plaid/plaid-go/plaid"
)

func ToAccounts(plaidAccounts []plaid.Account) []Account {
	accounts := make([]Account, len(plaidAccounts))

	for i, account := range plaidAccounts {
		accounts[i] = *ToAccount(&account)
	}

	return accounts
}

func ToAccount(plaidAccount *plaid.Account) *Account {
	return &Account{
		// Metadata
		PlaidId:      plaidAccount.AccountID,
		Mask:         plaidAccount.Mask,
		Name:         plaidAccount.Name,
		OfficialName: plaidAccount.OfficialName,
		Subtype:      plaidAccount.Subtype,
		Type:         plaidAccount.Type,

		// Balance
		AvailableBalance: plaidAccount.Balances.Available,
		CurrentBalance:   plaidAccount.Balances.Current,
		Limit:            plaidAccount.Balances.Limit,
		ISOCurrencyCode:  plaidAccount.Balances.ISOCurrencyCode,
	}
}

func FromUser(dbUser *User) *userpb.User {
	items := make([]*userpb.Item, len(dbUser.Items))
	for i, item := range dbUser.Items {
		items[i] = FromItem(&item)
	}

	return &userpb.User{
		Name:       dbUser.Name,
		PictureURL: dbUser.PictureURL,
		Email:      dbUser.Email,
		Id:         dbUser.AuthID,
		Items:      items,
	}
}

func FromItem(dbItem *Item) *userpb.Item {
	return &userpb.Item{
		Id:          dbItem.PlaidId,
		AccessToken: dbItem.PlaidAccessToken,
		Accounts:    FromAccounts(dbItem.Accounts),
	}
}

func FromAccounts(dbAccounts []Account) []*userpb.Account {
	accounts := make([]*userpb.Account, len(dbAccounts))
	for i, account := range dbAccounts {
		accounts[i] = FromAccount(&account)
	}

	return accounts
}

func FromAccount(dbAccount *Account) *userpb.Account {
	return &userpb.Account{
		Id:           dbAccount.PlaidId,
		Mask:         dbAccount.Mask,
		Name:         dbAccount.Name,
		OfficialName: dbAccount.OfficialName,
		Subtype:      dbAccount.Subtype,
		Type:         dbAccount.Type,

		AvailableBalance: dbAccount.AvailableBalance,
		CurrentBalance:   dbAccount.CurrentBalance,
		Limit:            dbAccount.Limit,
		Currency:         dbAccount.ISOCurrencyCode,
	}
}
