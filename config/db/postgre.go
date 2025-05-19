package db

import (
	"context"
	"dot-go/config/schema"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgre(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database :%v", err)
	}

	log.Println("connect to postgre")

	err = db.AutoMigrate(&schema.User{}, &schema.Music{}) //run once when the application starts
	if err != nil {
		log.Fatalf("failed to migrate schema :%v", err)
	}

	log.Println("success to migrate schema")

	err = SeedMusic(context.Background(), db) //seeder music
	if err != nil {
		log.Fatalf("failed to seed :%v", err)
	}

	return db
}

func SeedMusic(ctx context.Context, db *gorm.DB) error {

	var count int64
	db.Model(&schema.Music{}).Count(&count)
	if count > 0 {
		log.Println("Music already seeded")
		return nil
	}

	musics := []schema.Music{
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

	result := db.WithContext(ctx).Create(&musics)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("row affected :%d", result.RowsAffected)
	return nil
}
