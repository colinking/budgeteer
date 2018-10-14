package mock

import "github.com/colinking/budgeteer/backend/pkg/db"

type mockDatabase struct {
	user *db.User
}

var _ db.Database = (*mockDatabase)(nil)

// New generates a new Database.
func New() db.Database {
	return &mockDatabase{
		user: NewUser("Colin", "King"),
	}
}

func (d *mockDatabase) StoreUser(u *db.User) {
	d.user = u
}

func (d *mockDatabase) GetUser(userID string) *db.User {
	return d.user
}

func (d *mockDatabase) GetUserByEmail(email string) *db.User {
	return d.user
}

// SaveToken saves a token to the database for a given user.
func (d *mockDatabase) SaveToken(userID string, token string) {
	d.user.Plaid.AccessToken = token
}

func (d *mockDatabase) GetToken(userID string) string {
	return d.user.Plaid.AccessToken
}
