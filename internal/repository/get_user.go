package repository

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type GetUserQuery interface {
	GetUser(ctx context.Context, xid string) (*User, string)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type getUserQuery struct{}

func (d *getUserQuery) GetUser(ctx context.Context, xid string) (*User, string) {
	var user *User

	result := DB.Table("user").Select("*").Where("xid = $1", &xid).Scan(&user)
	if result.Error != nil {
		logrus.Error("error query get user: ", result.Error)
		return nil, result.Error.Error()
	} else if user == nil {
		logrus.Error("account not found")
		return nil, "not found"
	}

	logrus.Info("success query get account")
	return user, ""
}

func (d *getUserQuery) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user *User

	result := DB.Table("user").Select("*").Where("email = $1", &email).Scan(&user)
	if result.Error != nil {
		logrus.Error("error query get user: ", result.Error)
		return nil, result.Error
	} else if user == nil {
		logrus.Error("account not found")
		return nil, fmt.Errorf("not found")
	}

	logrus.Info("success query get account")
	return user, nil
}
