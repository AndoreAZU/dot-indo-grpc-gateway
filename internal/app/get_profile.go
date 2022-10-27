package app

import (
	"context"
	"dot-indonesia/internal/repository"
	"dot-indonesia/internal/util"
	desc "dot-indonesia/pkg"
	"encoding/json"
	"fmt"
	"time"

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

	// get from redis first
	var user repository.User
	if err = x.redis.GetKeyToStruct(ctx, xid, &user); err == nil {
		return &desc.ResponseGetProfile{
			Status:  200,
			Message: "success",
			Profile: util.UnmarshalUser(user),
		}, nil
	}

	user, err = x.getUser.GetUser(ctx, xid)
	if err != nil {
		return &desc.ResponseGetProfile{
			Status:  401,
			Message: err.Error(),
		}, nil
	}

	// store to redis
	a, _ := json.Marshal(user)
	err = x.redis.SetKey(ctx, xid, string(a), time.Duration(20*time.Minute))
	if err != nil {
		fmt.Println("error: ", err)
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
