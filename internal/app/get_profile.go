package app

import (
	"context"
	"dot-indonesia/internal/util"
	desc "dot-indonesia/pkg"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func (x *UserManagement) GetProfile(ctx context.Context, req *desc.EmptyMessage) (*desc.ResponseGetProfile, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md.Get("Authorization")[0]

	xid, err := x.redis.GetKeyString(ctx, token)
	if err != nil {
		fmt.Println(err)
		return &desc.ResponseGetProfile{
			Status:  401,
			Message: "session invalid",
		}, nil
	}

	user, err := x.getUser.GetUser(ctx, xid)
	if err != nil {
		return &desc.ResponseGetProfile{
			Status:  401,
			Message: err.Error(),
		}, nil
	}

	return &desc.ResponseGetProfile{
		Status:  200,
		Message: "success",
		Profile: util.UnmarshalUser(user),
	}, nil
}

func (x *UserManagement) GetUser(ctx context.Context, req *desc.User) (*desc.ResponseGetUser, error) {
	user, err := x.getUser.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &desc.ResponseGetUser{
			Status:  401,
			Message: err.Error(),
		}, nil
	}

	return &desc.ResponseGetUser{
		Status:  200,
		Message: "success",
		Profile: util.UnmarshalUserProfile(user),
	}, nil
}
