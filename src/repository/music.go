package repository

import (
	"context"
	"dot-go/src/model"
)

func (r *repository) GetMusics(ctx context.Context) ([]*model.Music, error) {

	var musics []*model.Music
	result := r.db.WithContext(ctx).Find(musics)
	if result.Error != nil {
		return nil, result.Error
	}

	return musics, nil

}
