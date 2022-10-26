package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"dot-indonesia/internal/repository"
	desc "dot-indonesia/pkg"
)

type Login interface {
	Login(ctx context.Context, param *desc.LoginRequest) (repository.User, error)
}

type login struct {
	dao repository.DAO
}

func NewLogin(dao repository.DAO) Login {
	return &login{dao: dao}
}

func (x *login) Login(ctx context.Context, param *desc.LoginRequest) (repository.User, error) {
	user, msg := x.dao.NewLoginQuery().Login(ctx, param.Email)
	if len(msg) != 0 {
		return repository.User{}, fmt.Errorf(msg)
	}

	// sha_pass := sha256.Sum256([]byte(param.Password))
	hasher := sha256.New()
	hasher.Write([]byte(param.Password))
	sha_pass := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(sha_pass, " ", user.Password)
	if strings.Compare(user.Password, sha_pass) != 0 {
		return repository.User{}, fmt.Errorf("invalid password")
	}

	return *user, nil
}
