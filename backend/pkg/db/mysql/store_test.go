package mysql

import (
	"testing"

	"github.com/colinking/budgeteer/backend/pkg/plaid"
	"github.com/stretchr/testify/require"

	"github.com/colinking/budgeteer/backend/pkg/db"
	_ "github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func setup(t *testing.T) *Database {
	// Boot an on-disk SQL database for testing.
	// d, err := gorm.Open("sqlite3", "/tmp/moss.store.test.db")

	// Boot an in-memory SQL database for testing.
	d, err := gorm.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("failed to boot gorm: %s", err)
	}

	// Clear previous table (if run against on-disk sqlite)
	d.DropTableIfExists("users")
	d.DropTableIfExists("items")
	d.DropTableIfExists("accounts")
	d.DropTableIfExists("institutions")

	Migrate(d)

	// Add some sample data.
	institution := &db.Institution{
		Name:      "institution-name-1",
		PlaidID:   "institution-id-1",
		BrandName: "institution-brand-name-1",
	}
	d.Save(institution)

	user := &db.User{
		AuthID: "1",
		Name:   "Colin King",
		Email:  "me@colinking.co",
		Items: []db.Item{
			{
				PlaidId:          "item-id-1",
				PlaidAccessToken: "item-token-1",
				Institution:      *institution,
				Accounts: []db.Account{
					{
						Name:    "account-name-1",
						PlaidId: "account-id-1",
					},
				},
			},
		},
	}
	d.Save(user)

	user = &db.User{
		AuthID: "10",
		Name:   "John Doe",
		Email:  "john@example.com",
	}
	d.Save(user)

	return &Database{
		db: d,
	}
}

func TestGetUser(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	user := d.GetUser(&db.GetUserInput{
		ID: "1",
	})
	require.NotNil(t, user)
	require.NotNil(t, user.User)
	require.Equal(t, "1", user.User.AuthID)
	require.Equal(t, "me@colinking.co", user.User.Email)
	require.Len(t, user.User.Items, 1)
	require.Len(t, user.User.Items[0].Accounts, 1)
	require.Equal(t, "account-id-1", user.User.Items[0].Accounts[0].PlaidId)
	require.Equal(t, "institution-brand-name-1", user.User.Items[0].Institution.BrandName)
}

func TestGetUserNonExistent(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	user := d.GetUser(&db.GetUserInput{
		ID: "non-existent",
	})
	require.NotNil(t, user)
	require.Nil(t, user.User)
}

func TestAddItem(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	user := d.GetUser(&db.GetUserInput{
		ID: "1",
	})
	require.Len(t, user.User.Items, 1)
	require.Equal(t, "item-id-1", user.User.Items[0].PlaidId)

	out := d.AddItem(&db.AddItemInput{
		AuthID:           "1",
		PlaidID:          "item-id-2",
		PlaidAccessToken: "item-token-2",
		InstitutionID:    "institution-id-1",
	})
	require.Equal(t, true, out.IsNew)

	user = d.GetUser(&db.GetUserInput{
		ID: "1",
	})
	require.Len(t, user.User.Items, 2)
	require.Equal(t, "item-id-1", user.User.Items[0].PlaidId)
	require.Equal(t, "item-id-2", user.User.Items[1].PlaidId)
	require.Equal(t, "institution-brand-name-1", user.User.Items[1].Institution.BrandName)
}

func TestUpsertUser(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	newUser := &db.UpsertUserInput{
		AuthID: "2",
		Name:   "Kyle King",
		Email:  "kyle@king.co",
	}
	require.Equal(t, true, d.UpsertUser(newUser).IsNew)

	updatedUser := &db.UpsertUserInput{
		AuthID: "2",
		Name:   "Kyle King",
		Email:  "new@email.com",
	}
	require.Equal(t, false, d.UpsertUser(updatedUser).IsNew)
}

func TestAddAccounts(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	out := d.AddAccounts(&db.AddAccountsInput{
		ItemID:   "item-id-1",
		Accounts: []db.Account{},
	})
	require.NotNil(t, out.User)
	require.Len(t, out.User.Items, 1)
	require.Len(t, out.User.Items[0].Accounts, 1)

	out = d.AddAccounts(&db.AddAccountsInput{
		ItemID: "item-id-1",
		Accounts: []db.Account{
			{
				Name:    "account-name-2",
				PlaidId: "account-id-2",
			},
			{
				Name:    "account-name-3",
				PlaidId: "account-id-3",
			},
		},
	})
	require.Len(t, out.User.Items, 1)
	require.Len(t, out.User.Items[0].Accounts, 3)
	account1 := out.User.Items[0].Accounts[0]
	account2 := out.User.Items[0].Accounts[1]
	account3 := out.User.Items[0].Accounts[2]
	require.Equal(t, "account-name-1", account1.Name)
	require.Equal(t, "account-name-2", account2.Name)
	require.Equal(t, "account-name-3", account3.Name)
	require.Equal(t, "account-id-1", account1.PlaidId)
	require.Equal(t, "account-id-2", account2.PlaidId)
	require.Equal(t, "account-id-3", account3.PlaidId)
}

func TestAddInstitution(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	institutionOut := d.AddInstitution(&db.AddInstitutionInput{
		Institution: &plaid.Institution{
			Name:      "institution-name-2",
			ID:        "institution-id-2",
			BrandName: "institution-brand-name-2",
		},
	})
	require.NotNil(t, institutionOut.Institution)
	require.Equal(t, "institution-name-2", institutionOut.Institution.Name)
	require.Equal(t, "institution-id-2", institutionOut.Institution.PlaidID)
}
