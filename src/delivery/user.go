package delivery

import (
	"dot-go/src/helper"
	"dot-go/src/helper/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *delivery) AddUserMusicFavorite(c echo.Context) error {

	var favoriteMusic validator.FavoriteMusic
	ctx, auth := d.getAuthAndCtx(c)

	if err := c.Bind(&favoriteMusic); err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := d.service.AddMusicFavoriteUser(ctx, auth, favoriteMusic.ID)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "Success", nil)

}
