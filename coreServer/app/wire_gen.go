// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"core-server/app/model/gormx/repo"
	"core-server/app/router"
	"core-server/app/service"
)

// Injectors from wire.go:

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	db, cleanup, err := InitGormDB()
	if err != nil {
		return nil, nil, err
	}
	demo := &repo.Demo{
		DB: db,
	}
	serviceDemo := &service.Demo{
		DemoModel: demo,
	}
	trans := &repo.Trans{
		DB: db,
	}
	user := &repo.User{
		DB: db,
	}
	serviceUser := &service.User{
		TransModel: trans,
		UserModel:  user,
	}
	routerRouter := &router.Router{
		DemoSrc: serviceDemo,
		UserSrc: serviceUser,
	}
	injector := &Injector{
		Router: routerRouter,
	}
	return injector, func() {
		cleanup()
	}, nil
}
