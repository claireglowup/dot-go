package service

import (
	"context"
	"dot-go/src/model"
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

	//task: mapping to make duration payload better (conver to menit and second)

	return payload, nil
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
