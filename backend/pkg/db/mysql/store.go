package mysql

import (
	"fmt"

	"github.com/colinking/budgeteer/backend/pkg/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/grpclog"
)

type Config struct {
	Port         int
	DatabaseName string
	Username     string
	Password     string
}

type Database struct {
	db *gorm.DB
}

func getInstitution(id string, database *gorm.DB) *db.Institution {
	institution := &db.Institution{
		PlaidID: id,
	}
	database.Set("gorm:auto_preload", true).First(institution)

	if institution.CreatedAt.IsZero() {
		return nil
	}

	return institution
}

func (d *Database) AddInstitution(input *db.AddInstitutionInput) *db.AddInstitutionOutput {
	tx := d.db.Begin()

	institution := &db.Institution{
		PlaidID:      input.Institution.ID,
		Name:         input.Institution.Name,
		BrandName:    input.Institution.BrandName,
		ColorDark:    input.Institution.Colors.Dark,
		ColorDarker:  input.Institution.Colors.Darker,
		ColorLight:   input.Institution.Colors.Light,
		ColorPrimary: input.Institution.Colors.Primary,
		Logo:         input.Institution.Logo,
		URL:          input.Institution.URL,
	}
	tx.Save(institution)

	tx.Commit()

	return &db.AddInstitutionOutput{
		Institution: institution,
	}
}

func (d *Database) AddAccounts(input *db.AddAccountsInput) *db.AddAccountsOutput {
	tx := d.db.Begin()

	item := getItem(input.ItemID, tx)

	// If the item doesn't exist, return a nil user.
	if item == nil {
		tx.Rollback()
		return &db.AddAccountsOutput{}
	}

	// Update the item with the new accounts.
	item.Accounts = append(item.Accounts, input.Accounts...)
	tx.Save(item)

	// Get the updated user.
	user := getUser(item.UserAuthID, tx)

	tx.Commit()

	return &db.AddAccountsOutput{
		User: user,
	}
}

func getItem(id string, database *gorm.DB) *db.Item {
	item := &db.Item{
		PlaidId: id,
	}
	database.Set("gorm:auto_preload", true).First(item)

	if item.CreatedAt.IsZero() {
		return nil
	}

	return item
}

func (d *Database) AddItem(input *db.AddItemInput) *db.AddItemOutput {
	tx := d.db.Begin()

	// Create the new Plaid Item.
	item := &db.Item{
		PlaidId:            input.PlaidID,
		PlaidAccessToken:   input.PlaidAccessToken,
		InstitutionPlaidID: input.InstitutionID,
	}
	tx.Save(item)

	// Add the Item to the current user.
	user := getUser(input.AuthID, tx)
	user.Items = append(user.Items, *item)
	tx.Save(user)

	tx.Commit()

	return &db.AddItemOutput{
		IsNew: true,
	}
}

func (d *Database) UpsertUser(input *db.UpsertUserInput) *db.UpsertUserOutput {
	user := &db.User{
		AuthID: input.AuthID,
	}
	count := 0

	tx := d.db.Begin()

	// Check if this user has logged in before.
	tx.Model(&user).Count(&count)
	isNew := count == 0

	// Upsert any new profile information.
	tx.First(&user)
	user.Email = input.Email
	user.Name = input.Name
	user.PictureURL = input.PictureURL

	tx.Save(&user)
	tx.Commit()

	return &db.UpsertUserOutput{
		IsNew: isNew,
		User:  user,
	}
}

func getUser(id string, database *gorm.DB) *db.User {
	user := &db.User{
		AuthID: id,
	}

	database.Set("gorm:auto_preload", true).First(user)

	if user.CreatedAt.IsZero() {
		return nil
	}

	return user
}

func (d *Database) GetUser(input *db.GetUserInput) *db.GetUserOutput {
	return &db.GetUserOutput{
		User: getUser(input.ID, d.db),
	}
}

// New initializes a new DynamoDB database connection.
func New(c *Config) (db.Database, error) {
	return openConnection(c)
}

func openConnection(c *Config) (db.Database, error) {
	endpoint := fmt.Sprintf("localhost:%d", c.Port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", c.Username, c.Password, endpoint, c.DatabaseName)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		return &Database{}, err
	}

	Migrate(d)

	grpclog.Infof("Opened MySQL connection: %s", endpoint)

	return &Database{
		db: d,
	}, nil
}
