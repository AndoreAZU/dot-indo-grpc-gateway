package repository

import (
	"context"
)

type RegisterQuery interface {
	Register(ctx context.Context, user User) error
}

type registerQuery struct{}

func (d *registerQuery) Register(ctx context.Context, user User) error {
	tx := DB.Begin()

	if err := tx.Table("user").Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
