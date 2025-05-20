package service

import (
	"context"
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
