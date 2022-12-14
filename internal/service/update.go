package service

import (
	"context"

	"dot-indonesia/internal/repository"
)

type Update interface {
	Update(ctx context.Context, xid string, user repository.User) error
}

type update struct {
	dao repository.DAO
}

func NewUpdate(dao repository.DAO) Update {
	return &login{dao: dao}
}

func (x *login) Update(ctx context.Context, xid string, user repository.User) error {
	err := x.dao.NewUpdateQuery().Update(ctx, xid, user)
	if err != nil {
		return err
	}

	return nil
}
