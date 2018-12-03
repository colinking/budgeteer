package mysql

import (
	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/jinzhu/gorm"
)

// https://github.com/go-gormigrate/gormigrate
func Migrate(d *gorm.DB) {
	d.AutoMigrate(
		db.User{},
		db.Item{},
	)

	// TODO: move to a better migration tool: https://github.com/go-gormigrate/gormigrate
	d.Model(db.User{}).RemoveIndex("email")
	d.Model(db.Item{}).RemoveIndex("plaid_id")
	d.Model(db.Item{}).RemoveIndex("plaid_access_token")
}
