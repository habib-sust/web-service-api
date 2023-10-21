//go:build wireinject
// +build wireinject

package config

import (
	"habib/web-service-api/app/controller"
	"habib/web-service-api/app/repository"
	"habib/web-service-api/app/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var albumServiceSet = wire.NewSet(
	service.AlbumServiceInit,
	wire.Bind(new(service.AlbumService), new(*service.AlbumServiceImpl)),
)

var albumRepositorySet = wire.NewSet(
	repository.AlbumRepositoryInit,
	wire.Bind(new(repository.AlbumRepository), new(*repository.AlbumRepositoryImpl)),
)

var albumControllerSet = wire.NewSet(
	controller.AlbumControllerInit,
	wire.Bind(new(controller.AlbumController), new(*controller.AlbumControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, albumControllerSet, albumServiceSet, albumRepositorySet)
	return nil
}
