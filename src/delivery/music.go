package delivery

import (
	"dot-go/src/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *delivery) GetMusics(c echo.Context) error {
	ctx, auth := d.getAuthAndCtx(c)

	result, err := d.service.GetMusic(ctx, auth)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "Success", result)
}
