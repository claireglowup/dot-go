package repository

import (
	"context"
	"dot-go/src/model"
)

func (r *repository) RegisterUser(ctx context.Context, user *model.User) error {

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
