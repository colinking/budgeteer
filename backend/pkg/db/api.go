package db

type UpsertUserInput struct {
	AuthID     string
	FirstName  string
	LastName   string
	Email      string
	PictureURL string
}

type UpsertUserOutput struct {
	IsNew bool
}

// Database stores a DB connection.
type Database interface {
	// User API
	GetUserByID(authID string) *User
	UpsertUser(input *UpsertUserInput) *UpsertUserOutput

	// User.Plaid API
	SaveToken(authID string, token string)
	GetToken(authID string) (token string)
}
