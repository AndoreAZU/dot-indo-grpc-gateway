package repository

import (
	"context"
	"fmt"
	"time"
)

type UpdateQuery interface {
	Update(ctx context.Context, xid string, user User) error
}

type updateQuery struct{}

func (d *updateQuery) Update(ctx context.Context, xid string, user User) error {

	tx := DB.Begin()
	tx = tx.Table("user").Where("xid = ?", xid).Updates(map[string]interface{}{
		"email":      user.Email,
		"username":   user.Email,
		"age":        user.Age,
		"address":    user.Address,
		"gender":     user.Gender,
		"updated_at": time.Now(),
	})

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}
