// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/data"
	"kratos-realworld/internal/server"
	"kratos-realworld/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	realWorldService := service.NewRealworldService(userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, realWorldService, logger)
	grpcServer := server.NewGRPCServer(confServer, realWorldService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
