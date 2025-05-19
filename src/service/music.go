package service

import (
	"context"
	"dot-go/src/model"
	"fmt"
)

func (s *service) GetMusic(ctx context.Context, authHeader string) (*[]model.Music, error) {

	_, err := s.getJWTClaims(authHeader)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.GetMusics(ctx)
	if err != nil {
		return nil, err
	}

	payload := make([]model.Music, 0, len(result))

	for _, music := range result {
		payload = append(payload, model.Music{
			Title:    music.Title,
			Artist:   music.Artist,
			Duration: fmt.Sprintf("%02d:%02d", music.Duration/60, music.Duration%60),
			Writer:   music.Writer,
			Year:     music.Year,
		})

	}

	return &payload, nil

}
