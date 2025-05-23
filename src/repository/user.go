package repository

import (
	"context"
	"dot-go/config/schema"
	"dot-go/src/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (r *repository) RegisterUser(ctx context.Context, user *schema.User) error {

	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	var user model.User

	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repository) AddMusicFavoriteUser(ctx context.Context, idUser uint, idMusic uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user schema.User
		var music schema.Music

		if err := tx.First(&user, idUser).Error; err != nil {
			return err
		}

		var existing []schema.Music
		if err := tx.Model(&user).
			Association("Favorites").
			Find(&existing, "id = ?", idMusic); err != nil {
			return err
		}

		if len(existing) > 0 {
			return gorm.ErrRecordNotFound
		}

		if err := tx.First(&music, idMusic).Error; err != nil {
			return err
		}

		if err := tx.Model(&user).Association("Favorites").Append(&music); err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) GetFavoriteMusicsByUser(ctx context.Context, idUser uint) (*[]schema.Music, error) {
	var user schema.User
	if err := r.db.WithContext(ctx).First(&user, idUser).Error; err != nil {
		return nil, err
	}

	var favorites []schema.Music
	if err := r.db.WithContext(ctx).Model(&user).Association("Favorites").Find(&favorites); err != nil {
		return nil, err
	}

	return &favorites, nil
}

func (r *repository) RemoveMusicFavoriteUser(ctx context.Context, idUser uint, idMusic uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user schema.User
		var music schema.Music

		if err := tx.First(&user, idUser).Error; err != nil {
			return err
		}

		if err := tx.First(&music, idMusic).Error; err != nil {
			return err
		}

		var existing []schema.Music
		if err := tx.Model(&user).
			Association("Favorites").
			Find(&existing, "id = ?", idMusic); err != nil {
			return err
		}

		if len(existing) == 0 {
			return errors.New(fmt.Sprintf("music with id: %d is not in your favorite", idMusic))
		}

		if err := tx.Model(&user).Association("Favorites").Delete(&music); err != nil {
			return err
		}

		return nil
	})
}
