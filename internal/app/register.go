package app

import (
	"context"
	"crypto/sha256"
	"dot-indonesia/internal/repository"
	desc "dot-indonesia/pkg"
	"encoding/hex"

	"github.com/google/uuid"
)

func (x *UserManagement) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.Response, error) {
	hasher := sha256.New()
	hasher.Write([]byte(req.Password))
	sha_pass := hex.EncodeToString(hasher.Sum(nil))
	user := repository.User{
		Xid:      uuid.NewString(),
		Email:    req.Email,
		Username: req.Username,
		Password: sha_pass,
		TierId:   "xid_1",
	}

	err := x.register.Register(ctx, user)
	if err != nil {
		return &desc.Response{
			Status:  401,
			Message: err.Error(),
			Data:    "",
		}, nil
	}

	return &desc.Response{
		Status:  200,
		Message: "success",
		Data:    user.Xid,
	}, nil
}
