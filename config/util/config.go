package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL       string
	HTTPServerAddress string
	SecretKey         string
}

func LoadConfig() (Config, error) {

	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		HTTPServerAddress: os.Getenv("HTTP_SERVER_ADDRESS"),
		SecretKey:         os.Getenv("SECRET_KEY"),
	}, nil

}
