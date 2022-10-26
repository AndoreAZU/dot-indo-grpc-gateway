package main

import (
	"context"
	"dot-indonesia/internal/app"
	"dot-indonesia/internal/repository"
	"dot-indonesia/internal/service"
	"dot-indonesia/internal/util"
	"log"
	"net/http"
	"os"

	desc "dot-indonesia/pkg"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
)

func main() {
	util.InitEnv()

	port_rest, port_grpc := "", ""
	if len(os.Getenv("PORT")) == 0 {
		port_rest = os.Getenv("REST_PORT")
	}

	// port_grpc := os.Getenv("GRPC_PORT")
	// port_rest := os.Getenv("REST_PORT")

	logrus.Info("starting port grpc: ", port_grpc, " port rest: ", port_rest)
	logrus.Info("env: " + os.Getenv("ENV"))

	db := repository.NewDBGorm()

	if db == nil {
		logrus.Error("failed initialize db")
		return
	}

	dao := repository.NewDAO(db)

	redis := service.NewRedis()
	login := service.NewLogin(dao)
	register := service.NewRegister(dao)
	getUser := service.NewGetUser(dao)
	update := service.NewUpdate(dao)
	follow := service.NewFollow(dao)

	app := app.UserManagementServer(
		login, redis, register, getUser, update, follow,
	)

	// disable grpc

	// go func() {
	// 	listener, err := net.Listen("tcp", port_grpc)
	// 	if err != nil {
	// 		logrus.Error("Error Start Server: ", err.Error())
	// 	}

	// 	grpcServer := grpc.NewServer()
	// 	desc.RegisterUserManagementServer(grpcServer, app)

	// 	err = grpcServer.Serve(listener)
	// 	if err != nil {
	// 		logrus.Error("Error Start Server: ", err.Error())
	// 	}
	// }()

	mux := runtime.NewServeMux()

	err := desc.RegisterUserManagementHandlerServer(context.Background(), mux, app)

	if err != nil {
		logrus.Error("Error Start Server: ", err.Error())
	} else {
		logrus.Info("service started")
	}

	log.Fatalln(http.ListenAndServe(port_rest, mux))
}
