package app

import (
	"context"
	"dot-indonesia/internal/util"
	desc "dot-indonesia/pkg"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func (x *UserManagement) Update(ctx context.Context, req *desc.User) (*desc.Response, error) {
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

	user_db, err := x.getUser.GetUser(ctx, xid)
	if err != nil {
		fmt.Println(err)
		return &desc.Response{
			Status:  401,
			Message: "session invalid",
		}, nil
	}

	user := util.MarshalUser(req)
	user.Xid = xid
	user.Password = user_db.Password

	err = x.update.Update(ctx, user)
	if err != nil {
		return &desc.Response{
			Status:  401,
			Message: err.Error(),
		}, nil
	}

	return &desc.Response{
		Status:  200,
		Message: "success",
	}, nil
}