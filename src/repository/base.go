package repository

import (
	"context"
	"dot-go/src/model"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetMusics(ctx context.Context) ([]*model.Music, error)
}

type repository struct {
	db *gorm.DB
}

func Newrepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
