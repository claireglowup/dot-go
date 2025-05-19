package service

import (
	"context"
	"dot-go/src/model"
	"fmt"
)

func (s *service) GetMusic(ctx context.Context) (*[]model.Music, error) {

	result, err := s.repo.GetMusics(ctx)
	if err != nil {
		return nil, err
	}

	payload := make([]model.Music, 0, len(result))

	for _, music := range result {
		payload = append(payload, model.Music{
			ID:       music.ID,
			Title:    music.Title,
			Artist:   music.Artist,
			Duration: fmt.Sprintf("%02d:%02d", music.Duration/60, music.Duration%60),
			Writer:   music.Writer,
			Year:     music.Year,
		})

	}

	return &payload, nil

}
