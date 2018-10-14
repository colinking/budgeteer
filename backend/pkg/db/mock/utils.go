package mock

import (
	"fmt"

	"github.com/colinking/budgeteer/backend/pkg/db"
)

// NewUser generates a new mock user.
func NewUser(firstName string, lastName string) *db.User {
	return &db.User{
		ID:        fmt.Sprintf("id-%s-%s", firstName, lastName),
		FirstName: firstName,
		LastName:  lastName,
		Email:     fmt.Sprintf("%s.%s@example.com", firstName, lastName),
		Plaid: &db.Plaid{
			AccessToken: fmt.Sprintf("token-%s-%s", firstName, lastName),
		},
	}
}
