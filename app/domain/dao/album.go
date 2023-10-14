package dao

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"->:false" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false" json:"-"`
}

// album represents data about a record album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	BaseModel
}
