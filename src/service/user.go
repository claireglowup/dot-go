package service

import (
	"context"
	"dot-go/src/model"
	"fmt"
	"strconv"
)

func (s *service) AddMusicFavoriteUser(ctx context.Context, authHeader string, idMusic uint) error {

	claims, err := s.getJWTClaims(authHeader)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return err
	}

	err = s.repo.AddMusicFavoriteUser(ctx, uint(id), idMusic)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetFavoriteMusicsByUser(ctx context.Context, authHeader string) (*[]model.Music, error) {
	claims, err := s.getJWTClaims(authHeader)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return nil, err
	}

	payload, err := s.repo.GetFavoriteMusicsByUser(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	payload2 := make([]model.Music, 0, len(*payload))

	for _, music := range *payload {
		payload2 = append(payload2, model.Music{
			ID:       music.ID,
			Title:    music.Title,
			Artist:   music.Artist,
			Duration: fmt.Sprintf("%02d:%02d", music.Duration/60, music.Duration%60),
			Writer:   music.Writer,
			Year:     music.Year,
		})

	}

	return &payload2, nil
}

func (s *service) UnfavoriteMusicUser(ctx context.Context, authHeader string, idMusic uint) error {
	claims, err := s.getJWTClaims(authHeader)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return err
	}

	err = s.repo.RemoveMusicFavoriteUser(ctx, uint(id), idMusic)
	if err != nil {
		return err
	}

	return nil
}
