package app

import (
	"dot-indonesia/internal/service"
	desc "dot-indonesia/pkg"
)

type UserManagement struct {
	desc.UnimplementedUserManagementServer
	login    service.Login
	redis    service.Redis
	register service.Register
	getUser  service.GetUser
	update   service.Update
	follow   service.Follow
}

func UserManagementServer(
	login service.Login, redis service.Redis, register service.Register, getUser service.GetUser, update service.Update,
	follow service.Follow,
) *UserManagement {
	return &UserManagement{login: login, redis: redis, register: register, getUser: getUser, update: update, follow: follow}
}
