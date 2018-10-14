package db

// Plaid stores metadata about a user's plaid account.
type Plaid struct {
	AccessToken string
}

// User stores metadata about one user account.
type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Plaid     *Plaid
}

// Database stores a DB connection.
type Database interface {
	// User API
	StoreUser(u *User)
	GetUser(userID string) *User
	GetUserByEmail(email string) *User

	// User.Plaid API
	SaveToken(userID string, token string)
	GetToken(userID string) (token string)
}
