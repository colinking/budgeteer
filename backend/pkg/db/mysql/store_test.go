package mysql

import (
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

	d.DropTableIfExists("users")
	d.DropTableIfExists("items")

	Migrate(d)

	// Add some sample data.
	user := &db.User{
		AuthID: "1",
		Name:   "Colin King",
		Email:  "me@colinking.co",
	}
	d.Save(user)

	return &Database{
		db: d,
	}
}

func TestAddItem(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	out := d.AddItem(&db.AddItemInput{
		AuthID:           "1",
		PlaidID:          "id-1",
		PlaidAccessToken: "token-1",
	})
	if out.IsNew != true {
		t.Fatalf("new item is not unique")
	}

	user := d.GetUserByID("1")
	if len(user.Items) != 1 {
		t.Fatalf("user not updated with new item")
	}

	if user.Items[0].PlaidId != "id-1" {
		t.Fatalf("user not updated with correct item")
	}
}

func TestUpsertUser(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	newUser := &db.UpsertUserInput{
		AuthID: "2",
		Name:   "Kyle King",
		Email:  "kyle@king.co",
	}

	if d.UpsertUser(newUser).IsNew != true {
		t.Fatalf("new user should be unique")
	}

	updatedUser := &db.UpsertUserInput{
		AuthID: "2",
		Name:   "Kyle King",
		Email:  "new@email.com",
	}

	if d.UpsertUser(updatedUser).IsNew != false {
		t.Fatalf("updated user should not be unique")
	}

	otherUser := &db.UpsertUserInput{
		AuthID: "3",
		Name:   "John Doe",
		Email:  "john@email.com",
	}

	if d.UpsertUser(otherUser).IsNew != true {
		t.Fatalf("other u ser should be unique")
	}
}
