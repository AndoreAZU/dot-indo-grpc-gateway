package service

import (
	"context"

	"dot-indonesia/internal/repository"
)

type Follow interface {
	Follow(ctx context.Context, xid, email string, action int32) error
	GetFollowing(ctx context.Context, xid string) ([]repository.User, error)
}

type follow struct {
	dao repository.DAO
}

func NewFollow(dao repository.DAO) Follow {
	return &follow{dao: dao}
}

func (x *follow) Follow(ctx context.Context, xid, email string, action int32) error {
	user, err := x.dao.NewGetUserQuery().GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	data := repository.Following{
		XidUser:      xid,
		XidFollowing: user.Xid,
	}

	if action == 1 {
		err = x.dao.NewFollowQuery().Follow(ctx, data)
	} else if action == 2 {
		err = x.dao.NewFollowQuery().Unfollow(ctx, data)
	}

	if err != nil {
		return err
	}

	return nil
}

func (x *follow) GetFollowing(ctx context.Context, xid string) ([]repository.User, error) {
	xid_following, err := x.dao.NewFollowQuery().GetFollowing(ctx, xid)
	users := make([]repository.User, len(xid_following))

	for i, v := range xid_following {
		u, _ := x.dao.NewGetUserQuery().GetUser(ctx, v.XidFollowing)
		users[i] = *u
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}
