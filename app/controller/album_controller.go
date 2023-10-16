package controller

import (
	"habib/web-service-api/app/service"

	"github.com/gin-gonic/gin"
)

type AlbumController interface {
	GetAllAlbums(c *gin.Context)
	AddNewAlbum(c *gin.Context)
	GetAlbumByID(c *gin.Context)
}

type AlbumControllerImpl struct {
	albumService service.AlbumService
}

func (controller AlbumControllerImpl) GetAllAlbums(c *gin.Context) {
	controller.albumService.GetAllAlbums(c)
}

func (controller AlbumControllerImpl) GetAlbumByID(c *gin.Context) {
	controller.albumService.GetAlbumById(c)
}

func (controller AlbumControllerImpl) AddNewAlbum(c *gin.Context) {
	controller.albumService.AddNewAlbum(c)
}

func AlbumControllerInit(albumService service.AlbumService) *AlbumControllerImpl {
	return &AlbumControllerImpl{albumService: albumService}
}
