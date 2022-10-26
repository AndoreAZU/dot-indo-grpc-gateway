package app

import (
	"context"
	desc "dot-indonesia/pkg"
	"time"

	"github.com/google/uuid"
)

func (x *UserManagement) Login(ctx context.Context, req *desc.LoginRequest) (*desc.Response, error) {
	user, err := x.login.Login(ctx, req)
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
