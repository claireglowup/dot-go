package repository

import (
	"context"
	"dot-go/config/schema"
)

func (r *repository) GetMusics(ctx context.Context) ([]schema.Music, error) {

	var musics []schema.Music

	result := r.db.WithContext(ctx).Find(&musics)
	if result.Error != nil {
		return nil, result.Error
	}

	return musics, nil

}
