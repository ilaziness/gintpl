package dao

import (
	"context"

	"github.com/ilaziness/gintpl/internal/ent"
	"github.com/ilaziness/gintpl/internal/ent/user"
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
	return client.User.Query().Where(user.ID(id)).Only(ctx)
}
func (*userType) List(ctx context.Context) (ent.Users, error) {
	return client.User.Query().All(ctx)
}
func (*userType) Create(ctx context.Context, user *ent.User) error {
	_, err := client.User.Create().
		SetUsername(user.Username).
		SetName(user.Name).
		SetAge(user.Age).
		Save(ctx)
	return err
}
