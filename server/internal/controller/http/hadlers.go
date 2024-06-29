package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/relationskatie/summer-practice/server/client"
)

func (ctrl *Controller) HandleGetHomePage(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandleGetVacancyByID(c echo.Context) error {
	id := c.Param("id")
	vacancyID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	vacancy, err := ctrl.store.Vacancies().GetVacancyById(c.Request().Context(), vacancyID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, vacancy)
}

func (ctrl *Controller) HandleGetForm(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandlePostForm(c echo.Context) error {
	return nil
}

func (ctrl *Controller) HandleGetAllVacancies(c echo.Context) error {
	model, _ := client.GetDataFromClient()
	return c.JSON(http.StatusOK, model)
}
