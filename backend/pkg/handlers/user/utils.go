package user

import (
	"fmt"
	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/colinking/budgeteer/backend/pkg/gen/userpb"
)

func ToUser(dbUser *db.User) *userpb.User {
	items := make([]*userpb.Item, len(dbUser.Items))
	for i, item := range dbUser.Items {
		items[i] = ToItem(&item)
	}

	return &userpb.User{
		Name:       dbUser.Name,
		PictureURL: dbUser.PictureURL,
		Email:      dbUser.Email,
		Id:         fmt.Sprint(dbUser.AuthID),
		Items:      items,
	}
}

func ToItem(dbItem *db.Item) *userpb.Item {
	return &userpb.Item{
		Id:          dbItem.PlaidId,
		AccessToken: dbItem.PlaidAccessToken,
	}
}
