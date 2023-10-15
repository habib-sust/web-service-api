package repository

import (
	"habib/web-service-api/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	GetAllAlbums() ([]dao.Album, error)
	GetAlbumById(id int) (dao.Album, error)
	AddNewAlbum(album dao.Album) ([]dao.Album, error)
}

type AlbumRepositoryImpl struct {
	db *gorm.DB
}

func (r AlbumRepositoryImpl) GetAllAlbums() ([]dao.Album, error) {
	var albums []dao.Album

	var err = r.db.Find(&albums).Error
	if err != nil {
		log.Error("Got an error finding all albums. Error: ", err)
		return nil, err
	}

	return albums, nil
}

func (r AlbumRepositoryImpl) GetAlbumById(id string) (dao.Album, error) {
	var album dao.Album

	err := r.db.Where("id = ?", id).First(&album).Error

	if err != nil {
		log.Error("Got an error when find album by id. Error: ", err)
		return dao.Album{}, err
	}

	return album, nil

}

func (r AlbumRepositoryImpl) AddNewAlbum(album *dao.Album) ([]dao.Album, error) {
	err := r.db.Save(album).Error
	if err != nil {
		log.Error("Got an error when save album. Error: ", err)
		return nil, err
	}

	var albums []dao.Album
	err = r.db.Find(&albums).Error
	if err != nil {
		log.Error("Got an error when finding all albums. Error: ", err)
		return nil, err
	}

	return albums, nil
}

func AlbumRepositoryInit(db *gorm.DB) *AlbumRepositoryImpl {
	db.AutoMigrate(&dao.Album{})

	return &AlbumRepositoryImpl{db: db}
}
