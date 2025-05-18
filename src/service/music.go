package service

import (
	"context"
	"dot-go/src/model"
)

func (s *service) GetMusic(ctx context.Context, authHeader string) ([]*model.Music, error) {

	_, err := s.getJWTClaims(authHeader)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.GetMusics(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil

}
