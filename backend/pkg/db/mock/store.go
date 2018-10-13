package mock

import "github.com/colinking/budgeteer/backend/pkg/db"

type mockDatabase struct {
	user *db.User
}

var _ db.Database = (*mockDatabase)(nil)

// New generates a new Database.
func New() db.Database {
	return &mockDatabase{
		user: db.NewUser("Colin King"),
	}
}

// SaveToken saves a token to the database for a given user.
func (d *mockDatabase) SaveToken(token string) {
	d.user.Plaid.AccessToken = token
}

func (d *mockDatabase) GetToken() string {
	return d.user.Plaid.AccessToken
}
