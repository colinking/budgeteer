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
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.PictureURL = input.PictureURL

	tx.Save(&user)
	tx.Commit()

	return &db.UpsertUserOutput{
		IsNew: isNew,
	}
}

func (d *Database) GetUserByID(authID string) *db.User {
	panic("implement me")
}

func (d *Database) SaveToken(authID string, token string) {
	panic("implement me")
}

func (d *Database) GetToken(authID string) (token string) {
	panic("implement me")
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