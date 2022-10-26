package repository

import (
	"context"

	"github.com/sirupsen/logrus"
)

type FollowQuery interface {
	Follow(ctx context.Context, data Following) error
	Unfollow(ctx context.Context, data Following) error
	GetFollowing(ctx context.Context, xid string) ([]Following, error)
	GetFollower(ctx context.Context, xid string) ([]User, error)
}

type followQuery struct{}

func (d *followQuery) Follow(ctx context.Context, data Following) error {
	tx := DB.Begin()

	err := tx.Table("following").Create(&data).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (d *followQuery) Unfollow(ctx context.Context, data Following) error {
	tx := DB.Begin()

	err := tx.Table("following").Delete(&data, "xid_user = ? and xid_following = ?", data.XidUser, data.XidFollowing).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (d *followQuery) GetFollowing(ctx context.Context, xid string) ([]Following, error) {
	var user *[]Following

	result := DB.Table("following").Select("xid_following").Where("xid_user = $1", &xid).Scan(&user)
	if result.Error != nil {
		logrus.Error("error query get user: ", result.Error)
		return nil, result.Error
	} else if user == nil {

		return []Following{}, nil
	}

	return *user, nil
}

func (d *followQuery) GetFollower(ctx context.Context, xid string) ([]User, error) {
	var user *[]User

	result := DB.Table("following").Select("xid_user").Where("xid = $1", &xid).Scan(&user)
	if result.Error != nil {
		logrus.Error("error query get user: ", result.Error)
		return nil, result.Error
	} else if user == nil {

		return []User{}, nil
	}

	return *user, nil
}
