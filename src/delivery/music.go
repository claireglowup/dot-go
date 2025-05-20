package delivery

import (
	"dot-go/src/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *delivery) getMusics(c echo.Context) error {

	result, err := d.service.GetMusic(c.Request().Context())
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "Success", result)
}
