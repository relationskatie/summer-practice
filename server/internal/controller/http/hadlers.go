package http

import (
	"io"

	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) HandleGetHomePage(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandleGetVacancyByID(c echo.Context) error {
	id := c.Param("id")
	if id == "123" {
		io.WriteString(c.Response(), "yees")
	}
	return nil
}

func (ctrl *Controller) HandleGetForm(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandleGetAllVacancies(c echo.Context) error {
	return nil
}
