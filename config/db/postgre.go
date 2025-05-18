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

	err = schema.SeedMusic(context.Background(), db) //seeder music
	if err != nil {
		log.Fatalf("failed to seed :%v", err)
	}

	return db
}

func CleanSeeders(db *gorm.DB) {

	if err := db.Migrator().DropTable(&schema.Music{}); err != nil {
		log.Fatalf("seeders not deleted :%v", err)
	}

	log.Println("seeders is deleted")

}
