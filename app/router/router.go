package router

import (
	"habib/web-service-api/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		album := api.Group("/album")
		album.GET("/albums", init.AlbumController.GetAllAlbums)
		album.GET("/:id", init.AlbumController.GetAlbumByID)
		album.POST("/album", init.AlbumController.AddNewAlbum)
	}

	return router
}
