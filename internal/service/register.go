package service

import (
	"context"

	"dot-indonesia/internal/repository"
)

type Register interface {
	Register(ctx context.Context, user repository.User) error
}

type register struct {
	dao repository.DAO
}

func NewRegister(dao repository.DAO) Register {
	return &register{dao: dao}
}

func (x *register) Register(ctx context.Context, user repository.User) error {
	err := x.dao.NewRegisterQuery().Register(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
