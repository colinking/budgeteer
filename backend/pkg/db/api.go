package db

// Plaid stores metadata about a user's plaid account.
type Plaid struct {
	accessToken string
}

// User stores metadata about one user account.
type User struct {
	name  string
	plaid *Plaid
}

func newUser(name string) *User {
	return &User{
		name:  name,
		plaid: &Plaid{},
	}
}

// Database stores a DB connection.
type Database interface {
	SaveToken(string)
	GetToken() string
}

type mockDatabase struct {
	user *User
}

// New generates a new Database.
func New() Database {
	return &mockDatabase{
		user: newUser("Colin King"),
	}
}

// SaveToken saves a token to the database for a given user.
func (db *mockDatabase) SaveToken(token string) {
	db.user.plaid.accessToken = token
}

func (db *mockDatabase) GetToken() string {
	return db.user.plaid.accessToken
}
