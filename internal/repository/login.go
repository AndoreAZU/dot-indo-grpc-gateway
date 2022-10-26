package repository

import (
	"context"

	"github.com/sirupsen/logrus"
)

type LoginQuery interface {
	Login(ctx context.Context, email string) (*User, string)
}

type loginQuery struct{}

func (d *loginQuery) Login(ctx context.Context, email string) (*User, string) {
	var user *User

	result := DB.Table("user").Select("*").Where("email = $1", &email).Scan(&user)
	if result.Error != nil {
		logrus.Error("error query login: ", result.Error)
		return nil, result.Error.Error()
	} else if user == nil {
		logrus.Error("account not found")
		return nil, "not found"
	}

	logrus.Info("success query get account")
	return user, ""
}
