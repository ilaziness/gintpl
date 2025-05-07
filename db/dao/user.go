package dao

import (
	"context"

	"github.com/ilaziness/gintpl/db"
	"github.com/ilaziness/gintpl/db/ent"
	"github.com/ilaziness/gintpl/db/ent/user"
)

var User IUser = &userType{}

type IUser interface {
	Get(ctx context.Context, id int) (*ent.User, error)
	List(ctx context.Context) (ent.Users, error)
	Create(ctx context.Context, user *ent.User) error
}

type userType struct {
}

func (*userType) Get(ctx context.Context, id int) (*ent.User, error) {
	return db.GetClient().User.Query().Where(user.ID(id)).Only(ctx)
}
func (*userType) List(ctx context.Context) (ent.Users, error) {
	return db.GetClient().User.Query().All(ctx)
}
func (*userType) Create(ctx context.Context, user *ent.User) error {
	_, err := db.GetClient().User.Create().
		SetUsername(user.Username).
		SetName(user.Name).
		SetAge(user.Age).
		Save(ctx)
	return err
}
