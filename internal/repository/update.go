package repository

import (
	"context"
)

type UpdateQuery interface {
	Update(ctx context.Context, user User) error
}

type updateQuery struct{}

func (d *updateQuery) Update(ctx context.Context, user User) error {

	tx := DB.Table("user").Save(user)
	if tx.Commit().Error != nil {
		tx.Rollback()
	}

	return nil
}
