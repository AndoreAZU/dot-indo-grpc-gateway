package app

import (
	"context"
	desc "dot-indonesia/pkg"

	"google.golang.org/grpc/metadata"
)

func (x *UserManagement) Logout(ctx context.Context, req *desc.EmptyMessage) (*desc.Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md.Get("Authorization")[0]

	err := x.redis.RemoveKey(ctx, token)

	if err != nil {
		return &desc.Response{
			Status:  401,
			Message: err.Error(),
			Data:    "",
		}, nil
	}
	return &desc.Response{
		Status:  200,
		Message: "success logout",
	}, nil
}
