package mysql

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/colinking/budgeteer/backend/pkg/db"
	_ "github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func setup(t *testing.T) *Database {
	// Boot an in-memory SQL database for testing.
	d, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to boot gorm: %s", err)
	}

	// Clear previous table (if run against on-disk sqlite)
	d.DropTableIfExists("users")
	d.DropTableIfExists("items")

	Migrate(d)

	// Add some sample data.
	user := &db.User{
		AuthID: "1",
		Name:   "Colin King",
		Email:  "me@colinking.co",
		Items: []db.Item{
			{
				PlaidId:          "plaid-id-1",
				PlaidAccessToken: "plaid-token-1",
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
	require.Equal(t, "plaid-id-1", user.User.Items[0].PlaidId)

	out := d.AddItem(&db.AddItemInput{
		AuthID:           "1",
		PlaidID:          "plaid-id-2",
		PlaidAccessToken: "plaid-token-2",
	})
	require.Equal(t, true, out.IsNew)

	user = d.GetUser(&db.GetUserInput{
		ID: "1",
	})
	require.Len(t, user.User.Items, 2)
	require.Equal(t, "plaid-id-1", user.User.Items[0].PlaidId)
	require.Equal(t, "plaid-id-2", user.User.Items[1].PlaidId)
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
