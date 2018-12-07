package db

import (
	"time"
)

//
// Conventions: http://doc.gorm.io/models.html
//

// Override gorm.Model to remove the `primary_key` tag
type Times struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// User stores metadata about one user account.
type User struct {
	Times

	// Auth0
	AuthID string `gorm:"not null;unique;primary_key"`

	// User Data
	Name       string
	Email      string `gorm:"not null"`
	PictureURL string

	// Plaid
	Items []Item
}

// Item stores metadata about a Plaid Item.
type Item struct {
	Times

	PlaidId          string `gorm:"not null;unique;primary_key"`
	PlaidAccessToken string `gorm:"not null;unique"`
	UserAuthID       string
}
