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
	params := model.FormResponse{
		Text:       request.Text,
		Salary:     request.Salary,
		Area:       request.Area,
		Employment: request.Employment,
		Experience: request.Experience,
	}
	ctrl.mutex.Lock()
	ctrl.data = append(ctrl.data, params)
	ctrl.mutex.Unlock()
	return c.JSON(http.StatusCreated, params)
}

func (ctrl *Controller) HandleGetAllVacancies(c echo.Context) error {
	ctrl.mutex.Lock()
	data := ctrl.data
	ctrl.data = []model.FormResponse{}
	ctrl.mutex.Unlock()
	modelType := data[0]
	mode, err := client.GetDataFromClient(modelType.Text, modelType.Salary, modelType.Area, modelType.Employment, modelType.Experience)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, mode)
}
