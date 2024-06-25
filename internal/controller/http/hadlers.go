package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) HadleHi(c echo.Context) error {
	return c.String(http.StatusOK, "hi")
}
