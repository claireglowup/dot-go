package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string  `gorm:"type:varchar(100)"`
	Email     string  `gorm:"type:varchar(150);unique"`
	Password  string  `gorm:"type:varchar(255)"`
	Favorites []Music `gorm:"many2many:user_music_favorites"`
}

type Music struct {
	gorm.Model
	Title    string `gorm:"type:varchar(100);not null"`
	Artist   string `gorm:"type:varchar(100);not null"`
	Duration int    `gorm:"not null"`
	Writer   string `gorm:"type:varchar(100);default:'-'" `
	Year     int    `gorm:"not null;check:year >= 1900"`
}
