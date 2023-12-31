// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"github.com/google/wire"
	"habib/web-service-api/app/controller"
	"habib/web-service-api/app/repository"
	"habib/web-service-api/app/service"
)

// Injectors from wire.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	albumRepositoryImpl := repository.AlbumRepositoryInit(gormDB)
	albumServiceImpl := service.AlbumServiceInit(albumRepositoryImpl)
	albumControllerImpl := controller.AlbumControllerInit(albumServiceImpl)
	initialization := NewInitialization(albumRepositoryImpl, albumServiceImpl, albumControllerImpl)
	return initialization
}

// wire.go:

var db = wire.NewSet(ConnectToDB)

var albumServiceSet = wire.NewSet(service.AlbumServiceInit, wire.Bind(new(service.AlbumService), new(*service.AlbumServiceImpl)))

var albumRepositorySet = wire.NewSet(repository.AlbumRepositoryInit, wire.Bind(new(repository.AlbumRepository), new(*repository.AlbumRepositoryImpl)))

var albumControllerSet = wire.NewSet(controller.AlbumControllerInit, wire.Bind(new(controller.AlbumController), new(*controller.AlbumControllerImpl)))
