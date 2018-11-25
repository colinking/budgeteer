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

	db.Migrate(d)

	return &Database{
		db: d,
	}
}

func TestUpsertUser(t *testing.T) {
	d := setup(t)
	defer d.db.Close()

	newUser := &db.UpsertUserInput{
		AuthID:    "1",
		FirstName: "Colin",
		LastName:  "King",
		Email:     "me@colinking.co",
	}

	if d.UpsertUser(newUser).IsNew != true {
		t.Errorf("new user should be unique")
	}

	updatedUser := &db.UpsertUserInput{
		AuthID:    "1",
		FirstName: "Colin",
		LastName:  "King",
		Email:     "new@email.com",
	}

	if d.UpsertUser(updatedUser).IsNew != false {
		t.Errorf("updated user should not be unique")
	}

	otherUser := &db.UpsertUserInput{
		AuthID:    "2",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@email.com",
	}

	if d.UpsertUser(otherUser).IsNew != true {
		t.Errorf("other user should be unique")
	}
}

//func TestSaveToken(t *testing.T) {
//	d := setup(t)
//	token := "hello world"
//	d.SaveToken("1234", token)
//
//	if d.GetToken("1234") != token {
//		t.Errorf("Invalid token returned")
//	}
//}
