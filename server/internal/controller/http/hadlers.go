package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/relationskatie/summer-practice/server/client"
	"github.com/relationskatie/summer-practice/server/internal/model"
	"go.uber.org/zap"
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
	ctrl.log.Info("parce vacancy")
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
		Text:   request.Text,
		Salary: request.Salary,
		Area:   request.Area,
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

	if len(data) == 0 {
		return c.JSON(http.StatusNotFound, "no data available")
	}

	modelType := data[0]
	mode, err := client.GetDataFromClient(modelType.Text, modelType.Salary, modelType.Area)
	if err != nil {
		ctrl.log.Error("Error getting data from client", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = ctrl.store.Vacancies().DeleteAll(c.Request().Context())
	if err != nil {
		ctrl.log.Info("No vacancies found in DB to delete", zap.Error(err))
	}
	ctrl.log.Info("Delete all old vacancies")
	err = ctrl.store.Vacancies().AppendAll(c.Request().Context(), mode)
	if err != nil {
		ctrl.log.Error("Error appending all vacancies", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err)
	}
	ctrl.log.Info("Add Vacancies in db")

	vacancies, err := ctrl.store.Vacancies().GetAll(c.Request().Context())
	if err != nil {
		ctrl.log.Error("Error getting all vacancies", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err)
	}
	ctrl.log.Info("Get all vacancies")

	return c.JSON(http.StatusOK, vacancies)
}
