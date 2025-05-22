package delivery

import (
	"dot-go/src/helper"
	"dot-go/src/helper/validator"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (d *delivery) addUserMusicFavorite(c echo.Context) error {

	var favoriteMusic validator.FavoriteMusic
	ctx, auth := d.getAuthAndCtx(c)

	if err := c.Bind(&favoriteMusic); err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := c.Validate(favoriteMusic)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = d.service.AddMusicFavoriteUser(ctx, auth, favoriteMusic.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.WriteResponse(c, http.StatusNotFound, fmt.Sprintf("music with id %d is not found", favoriteMusic.ID), nil)
		}
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusOK, "Success", nil)

}

func (d *delivery) getFavoriteMusicsByUser(c echo.Context) error {

	ctx, auth := d.getAuthAndCtx(c)

	payload, err := d.service.GetFavoriteMusicsByUser(ctx, auth)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	if len(*payload) == 0 {
		return helper.WriteResponse(c, http.StatusOK, "Success", "you haven't added your favorite music")
	}

	return helper.WriteResponse(c, http.StatusOK, "Success", payload)
}
