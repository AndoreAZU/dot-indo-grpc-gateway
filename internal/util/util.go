package util

import (
	"dot-indonesia/internal/repository"
	"os"
	"time"

	"path/filepath"
	"runtime"

	desc "dot-indonesia/pkg"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	_, b, _, _      = runtime.Caller(0)
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../../")
)

func InitEnv() {
	profile := os.Getenv("ENV")
	logrus.Info("start load config profile ", profile)
	err := godotenv.Load(ProjectRootPath + "/resource/env")
	if err != nil {
		logrus.Error("failed init env from file : ", err.Error())
	}

	// load config profile from file env
	err = godotenv.Load(ProjectRootPath + "/resource/env-" + profile)
	if err != nil {
		logrus.Error("failed init env from file resource/env-: ", profile, err.Error())
	}
}

func UnmarshalUser(u repository.User) *desc.User {
	return &desc.User{
		Xid:       u.Xid,
		Email:     u.Email,
		Username:  u.Username,
		Age:       int32(u.Age),
		Address:   u.Address,
		Gender:    u.Gender,
		TierId:    u.TierId,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}

func MarshalUser(u *desc.User) repository.User {
	c, _ := time.Parse(time.RFC3339, u.CreatedAt)
	d, _ := time.Parse(time.RFC3339, u.UpdatedAt)

	return repository.User{
		Xid:       u.Xid,
		Email:     u.Email,
		Username:  u.Username,
		Age:       int(u.Age),
		Address:   u.Address,
		Gender:    u.Gender,
		TierId:    u.TierId,
		CreatedAt: c,
		UpdatedAt: d,
	}
}

func UnmarshalUserProfile(u repository.User) *desc.UserProfile {
	return &desc.UserProfile{
		Xid:      u.Xid,
		Username: u.Username,
		Email:    u.Email,
	}
}
