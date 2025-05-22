package service

import (
	"context"
	"dot-go/config/schema"
	"dot-go/src/helper/validator"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUser(ctx context.Context, user validator.UserRegister) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return err
	}
	payload := &schema.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(passwordHash),
	}

	err = s.repo.RegisterUser(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *service) Login(ctx context.Context, user validator.UserLogin) (string, error) {

	data, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", errors.New("email is not found, register please")
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("wrong password")

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    data.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil

}
