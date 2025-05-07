package service

import (
	"context"

	"github.com/ilaziness/gintpl/app/web/api"
	"github.com/ilaziness/gintpl/db/dao"
	"github.com/ilaziness/gintpl/db/ent"
	"github.com/ilaziness/gokit/log"
)

var User = &userService{}

type userService struct {
}

func (userService) Create(ctx context.Context, req *api.UserCreateReq) error {
	err := dao.User.Create(ctx, &ent.User{
		Age:      req.Age,
		Name:     req.Name,
		Username: req.Username,
	})

	return err
}

func (userService) Get(ctx context.Context, req *api.UserGetReq) (*api.UserGetRes, error) {
	user, err := dao.User.Get(ctx, req.ID)
	if err != nil {
		log.Error(ctx, "get user failed, err: %v", err)
		return nil, err
	}

	return &api.UserGetRes{
		Age:      user.Age,
		Name:     user.Name,
		Username: user.Username,
	}, nil
}

func (userService) List(_ context.Context) error {
	return nil
}

func (userService) List2(ctx context.Context) ([]*api.UserGetRes, error) {
	user, _ := dao.User.Get(ctx, 1)
	return []*api.UserGetRes{
		{
			Age:      user.Age,
			Name:     user.Name,
			Username: user.Username,
		},
	}, nil
}
