package service

import (
	"context"
	"dot-go/src/helper/validator"
	"dot-go/src/model"
	"dot-go/src/repository"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	Login(ctx context.Context, user validator.UserLogin) (string, error)
	RegisterUser(ctx context.Context, user validator.UserRegister) error
	GetMusic(ctx context.Context) (*[]model.Music, error)
	AddMusicFavoriteUser(ctx context.Context, authHeader string, idMusic uint) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) getJWTClaims(authHeader string) (*jwt.RegisteredClaims, error) {

	parts := strings.Split(authHeader, " ")

	tokenString := parts[1]

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(*jwt.RegisteredClaims), nil

}
