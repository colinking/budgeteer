package db

// Plaid stores metadata about a user's plaid account.
type Plaid struct {
	AccessToken string
}

// User stores metadata about one user account.
type User struct {
	Name  string
	Plaid *Plaid
}

// Database stores a DB connection.
type Database interface {
	SaveToken(string)
	GetToken() string
}
