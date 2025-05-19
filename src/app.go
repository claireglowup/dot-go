package src

import (
	"dot-go/config/db"
	"dot-go/config/util"
	"dot-go/src/delivery"
	"dot-go/src/helper"
	"dot-go/src/helper/validator"
	"dot-go/src/repository"
	"dot-go/src/service"
	"log"
	"net/http"

	validatorEngine "github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server interface {
	Run()
	CleanSeeders()
}

type server struct {
	httpServer *echo.Echo
	config     util.Config
	db         *gorm.DB
}

func InitServer() Server {

	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load env:", err)
	}
	postgre := db.InitPostgre(config.DatabaseURL)

	e := echo.New()
	e.Validator = &validator.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return &server{
		config:     config,
		httpServer: e,
		db:         postgre,
	}

}

func (s *server) Run() {

	repo := repository.Newrepository(s.db)
	service := service.NewService(repo)
	delivery := delivery.NewDelivery(service)

	configJWT := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey: []byte(s.config.SecretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
		},
	}

	delivery.Routes(s.httpServer, configJWT)

	if err := s.httpServer.Start(s.config.HTTPServerAddress); err != nil {
		log.Fatal(err.Error())
	}

}

func (s *server) CleanSeeders() {
	db.CleanSeeders(s.db)
}
