package app

import (
	"context"
	"dot-indonesia/internal/util"
	desc "dot-indonesia/pkg"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func (x *UserManagement) Follow(ctx context.Context, req *desc.FollowRequest) (*desc.Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md.Get("Authorization")[0]

	xid, err := x.redis.GetKeyString(ctx, token)
	if err != nil {
		fmt.Println(err)
		return &desc.Response{
			Status:  401,
			Message: "session invalid",
		}, nil
	}

	user, _ := x.getUser.GetUser(ctx, xid)

	// 1 follow 2 unfollow
	err = x.follow.Follow(ctx, user.Xid, req.Email, req.Action)
	if err != nil {
		return &desc.Response{
			Status:  401,
			Message: err.Error(),
			Data:    "",
		}, nil
	}

	session := uuid.New().String()
	x.redis.SetKey(ctx, session, user.Xid, time.Duration(20*time.Minute))

	return &desc.Response{
		Status:  200,
		Message: "success",
		Data:    session,
	}, nil
}

func (x *UserManagement) GetFollowing(ctx context.Context, req *desc.EmptyMessage) (*desc.ResponseGetFollowing, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md.Get("Authorization")[0]

	xid, err := x.redis.GetKeyString(ctx, token)
	if err != nil {
		return &desc.ResponseGetFollowing{
			Status:  500,
			Message: "failed",
		}, nil
	}

	following, err := x.follow.GetFollowing(ctx, xid)

	if err != nil {
		return &desc.ResponseGetFollowing{
			Status:  401,
			Message: err.Error(),
		}, nil
	}

	users := make([]*desc.UserProfile, len(following))
	for i, v := range following {
		users[i] = util.UnmarshalUserProfile(v)
	}

	return &desc.ResponseGetFollowing{
		Status:  200,
		Message: "success",
		User:    users,
	}, nil
}
