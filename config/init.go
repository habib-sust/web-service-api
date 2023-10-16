package config

import (
	"habib/web-service-api/app/controller"
	"habib/web-service-api/app/repository"
	"habib/web-service-api/app/service"
)

type Initialization struct {
	albumRepository repository.AlbumRepository
	albumService    service.AlbumService
	AlbumController controller.AlbumController
}

func NewInitialization(
	albumRepository repository.AlbumRepository,
	albumService service.AlbumService,
	albumController controller.AlbumController,
) *Initialization {
	return &Initialization{
		albumRepository: albumRepository,
		albumService:    albumService,
		AlbumController: albumController,
	}
}
