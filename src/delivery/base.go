package delivery

import (
	"context"
	"dot-go/src/service"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Delivery interface {
	Routes(e *echo.Echo, configJWT echojwt.Config)
}

type delivery struct {
	service service.Service
}

func NewDelivery(service service.Service) Delivery {
	return &delivery{
		service: service,
	}
}

func (d *delivery) getAuthAndCtx(c echo.Context) (context.Context, string) {
	return c.Request().Context(), c.Request().Header.Get("Authorization")

}

func (d *delivery) auth(e *echo.Group) {

	e.POST("/register", d.registerHandler)
	e.POST("/login", d.login)
}

func (d *delivery) music(e *echo.Group, configJwt echojwt.Config) {
	e.Use(echojwt.WithConfig(configJwt))
	e.GET("/musics", d.getMusics)
}

func (d *delivery) user(e *echo.Group, configJwt echojwt.Config) {
	e.Use(echojwt.WithConfig(configJwt))
	e.POST("/favorite", d.addUserMusicFavorite)
	e.GET("/favorite", d.getFavoriteMusicsByUser)
}

func (d *delivery) Routes(e *echo.Echo, configJWT echojwt.Config) {
	d.auth(e.Group("auth"))
	d.music(e.Group("music"), configJWT)
	d.user(e.Group("user"), configJWT)

}
