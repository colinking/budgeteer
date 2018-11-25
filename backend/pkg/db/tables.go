package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

//
// Conventions: http://doc.gorm.io/models.html
//

// Override gorm.Model to remove the `primary_key` tag
type GormModelCustomPrimaryKey struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// User stores metadata about one user account.
type User struct {
	GormModelCustomPrimaryKey

	// Auth0
	AuthID string `gorm:"not null;unique;primary_key"`

	// User Data
	FirstName  string
	LastName   string
	Email      string `gorm:"not null"`
	PictureURL string

	// Plaid
	Items []Item
}

type Item struct {
	gorm.Model
	PlaidId          string `gorm:"not null;unique"`
	PlaidAccessToken string `gorm:"not null;unique"`
}
