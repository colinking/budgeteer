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

	// Metadata
	PlaidId          string `gorm:"not null;unique;primary_key"`
	PlaidAccessToken string `gorm:"not null;unique"`

	// Accounts
	Accounts    []Account
	Institution Institution

	// Foreign Keys
	UserAuthID         string `gorm:"not null"`
	InstitutionPlaidID string `gorm:"not null"`
}

// Account stores metadata about a Plaid Account.
type Account struct {
	Times

	// Metadata
	PlaidId      string `gorm:"not null;unique;primary_key"`
	Mask         string
	Name         string
	OfficialName string
	Subtype      string
	Type         string

	// Balance
	AvailableBalance float64
	CurrentBalance   float64
	Limit            float64
	ISOCurrencyCode  string

	// Foreign Keys
	ItemPlaidID string
}

// Institution stores metadata about a Plaid Institution.
type Institution struct {
	Times

	// Metadata
	PlaidID   string `gorm:"not null;unique;primary_key"`
	BrandName string
	Name      string
	Logo      string `gorm:"type:varchar(4000)"`
	URL       string

	// Colors
	ColorDark    string
	ColorDarker  string
	ColorLight   string
	ColorPrimary string
}
