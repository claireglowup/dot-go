package delivery

import (
	"dot-go/src/helper"
	"dot-go/src/helper/validator"
	"fmt"
	"net/http"
	"strconv"

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
		if err == gorm.ErrDuplicatedKey {
			return helper.WriteResponse(c, http.StatusBadRequest, fmt.Sprintf("music with id %d is already favorited", favoriteMusic.ID), nil)
		}
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusCreated, "Success", nil)

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

func (d *delivery) UnfavoriteMusicUser(c echo.Context) error {

	ctx, auth := d.getAuthAndCtx(c)
	id := c.Param("id_music")

	idMusic, err := strconv.Atoi(id)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	err = d.service.UnfavoriteMusicUser(ctx, auth, uint(idMusic))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.WriteResponse(c, http.StatusNotFound, fmt.Sprintf("music with id %d is not found in your favorite", idMusic), nil)
		}
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusOK, "Success", nil)

}
