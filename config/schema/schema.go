package schema

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(150);unique"`
	Password string `gorm:"type:varchar(255)"`
}

type Music struct {
	gorm.Model
	Title    string `gorm:"type:varchar(100);not null"`
	Artist   string `gorm:"type:varchar(100);not null"`
	Duration int    `gorm:"not null"`
	Writer   string `gorm:"type:varchar(100);default:'-'"`
	Year     int    `gorm:"not null;check:year >= 1900"`
}

func SeedMusic(ctx context.Context, db *gorm.DB) error {
	musics := []*Music{
		{Title: "Love Story", Artist: "Taylor Swift", Duration: 230, Writer: "Taylor Swift", Year: 2008},
		{Title: "You Belong with Me", Artist: "Taylor Swift", Duration: 215, Writer: "Taylor Swift", Year: 2008},
		{Title: "Blank Space", Artist: "Taylor Swift", Duration: 231, Writer: "Taylor Swift", Year: 2014},
		{Title: "Shake It Off", Artist: "Taylor Swift", Duration: 242, Writer: "Taylor Swift", Year: 2014},
		{Title: "Wildest Dreams", Artist: "Taylor Swift", Duration: 220, Writer: "Taylor Swift", Year: 2014},
		{Title: "Look What You Made Me Do", Artist: "Taylor Swift", Duration: 211, Writer: "Taylor Swift", Year: 2017},
		{Title: "Lover", Artist: "Taylor Swift", Duration: 221, Writer: "Taylor Swift", Year: 2019},
		{Title: "Cardigan", Artist: "Taylor Swift", Duration: 239, Writer: "Taylor Swift", Year: 2020},
		{Title: "Willow", Artist: "Taylor Swift", Duration: 214, Writer: "Taylor Swift", Year: 2020},
		{Title: "Anti-Hero", Artist: "Taylor Swift", Duration: 200, Writer: "Taylor Swift", Year: 2022},
	}

	result := db.WithContext(ctx).Create(musics)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("row affected :%d", result.RowsAffected)
	return nil
}
