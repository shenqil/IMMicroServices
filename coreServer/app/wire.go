//go:build wireinject
// +build wireinject

package app

import (
	"core-server/app/model/gormx/repo"
	"core-server/app/router"
	"core-server/app/service"
	"github.com/google/wire"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		repo.RepoSet,
		service.ServiceSet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
