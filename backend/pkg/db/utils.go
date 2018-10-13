package db

// NewUser generates a new user.
func NewUser(name string) *User {
	return &User{
		Name:  name,
		Plaid: &Plaid{},
	}
}
