package repository

import (
	"context"
	"dot-go/config/schema"
	"dot-go/src/model"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(ctx context.Context, user *schema.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetMusics(ctx context.Context) ([]schema.Music, error)
	AddMusicFavoriteUser(ctx context.Context, idUser uint, idMusic uint) error
	GetFavoriteMusicsByUser(ctx context.Context, idUser uint) (*[]model.Music, error)
	RemoveMusicFavoriteUser(ctx context.Context, idUser uint, idMusic uint) error
}

type repository struct {
	db *gorm.DB
}

func Newrepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
