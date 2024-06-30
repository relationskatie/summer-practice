package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/relationskatie/summer-practice/server/client"
	"github.com/relationskatie/summer-practice/server/internal/model"
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
	var (
		request model.FormRequest
	)
	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest, err)
	}
	params := &model.FormResponse{
		ID:         uuid.New(),
		Text:       request.Text,
		Salary:     request.Salary,
		Area:       request.Area,
		URL:        request.URL,
		Employment: request.Employment,
		Experience: request.Experience,
	}
	return c.JSON(http.StatusCreated, params)

}

func (ctrl *Controller) HandleGetAllVacancies(c echo.Context) error {
	model, _ := client.GetDataFromClient()
	return c.JSON(http.StatusOK, model)
}
