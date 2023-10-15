package service

import (
	"habib/web-service-api/app/constant"
	"habib/web-service-api/app/domain/dao"
	"habib/web-service-api/app/pkg"
	"habib/web-service-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AlbumService interface {
	GetAllAlbums(c *gin.Context)
	GetAlbumById(c *gin.Context)
	AddNewAlbum(c *gin.Context)
}

type AlbumServiceImpl struct {
	albumRepository repository.AlbumRepository
}

func (s AlbumServiceImpl) GetAllAlbums(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all albums")

	data, err := s.albumRepository.GetAllAlbums()

	if err != nil {
		log.Error("Error in AlbumService when find all albums. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.IndentedJSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s AlbumServiceImpl) GetAlbumById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get album by id")
	id := c.Param("id")
	data, err := s.albumRepository.GetAlbumById(id)

	if err != nil {
		log.Error("Error in AlbumService when finding album by id")
		pkg.PanicException(constant.DataNotFound)
	}

	c.IndentedJSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s AlbumServiceImpl) AddNewAlbum(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute add new album")

	var album dao.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		log.Error("Error in AlbumService when mapping request from FE. Error: ", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.albumRepository.AddNewAlbum(&album)

	if err != nil {
		log.Error("Error in AlbumService when saving data to database. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.IndentedJSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func AlbumServiceInit(albumRepository repository.AlbumRepository) *AlbumServiceImpl {
	return &AlbumServiceImpl{albumRepository: albumRepository}
}
