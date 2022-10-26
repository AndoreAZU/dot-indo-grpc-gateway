package service

import (
	"context"
	"fmt"

	"dot-indonesia/internal/repository"
)

type GetUser interface {
	GetUser(ctx context.Context, xid string) (repository.User, error)
	GetUserByEmail(ctx context.Context, email string) (repository.User, error)
}

type getUser struct {
	dao repository.DAO
}

func NewGetUser(dao repository.DAO) GetUser {
	return &getUser{dao: dao}
}

func (x *getUser) GetUser(ctx context.Context, xid string) (repository.User, error) {
	user, msg := x.dao.NewGetUserQuery().GetUser(ctx, xid)
	if len(msg) != 0 {
		return repository.User{}, fmt.Errorf(msg)
	}
	return *user, nil
}

func (x *getUser) GetUserByEmail(ctx context.Context, email string) (repository.User, error) {
	user, err := x.dao.NewGetUserQuery().GetUserByEmail(ctx, email)
	if err != nil {
		return repository.User{}, err
	}
	return *user, nil
}
