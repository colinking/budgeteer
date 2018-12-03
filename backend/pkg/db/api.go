package db

type UpsertUserInput struct {
	AuthID     string
	Name       string
	Email      string
	PictureURL string
}

type UpsertUserOutput struct {
	IsNew bool
}

type AddItemInput struct {
	AuthID           string
	PlaidID          string
	PlaidAccessToken string
}

type AddItemOutput struct {
	IsNew bool
}

// Database stores a DB connection.
type Database interface {
	// User API
	GetUserByID(authID string) *User
	UpsertUser(input *UpsertUserInput) *UpsertUserOutput

	AddItem(input *AddItemInput) *AddItemOutput
}
