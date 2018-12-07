package db

// Database stores a DB connection.
type Database interface {
	// User API
	GetUser(input *GetUserInput) *GetUserOutput
	UpsertUser(input *UpsertUserInput) *UpsertUserOutput
	AddItem(input *AddItemInput) *AddItemOutput
	AddAccounts(input *AddAccountsInput) *AddAccountsOutput
}

type UpsertUserInput struct {
	AuthID     string
	Name       string
	Email      string
	PictureURL string
}

type UpsertUserOutput struct {
	IsNew bool
	User  *User
}

type AddItemInput struct {
	AuthID           string
	PlaidID          string
	PlaidAccessToken string
}

type AddItemOutput struct {
	IsNew bool
}

type GetUserInput struct {
	ID string
}

type GetUserOutput struct {
	User *User
}

type AddAccountsInput struct {
	ItemID   string
	Accounts []Account
}

type AddAccountsOutput struct {
	User *User
}
