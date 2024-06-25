package http

import (
	"io"

	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) HandleGetVacanciesByTunning(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandleGetVacancyById(c echo.Context) error {
	id := c.Param("id")
	if id == "123" {
		io.WriteString(c.Response(), "yees")
	}
	return nil
}

func (ctrl *Controller) HandleAddToFavourite(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandleDeleteIntoFavourite(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandleGetAllFavourite(c echo.Context) error {
	return nil
}
