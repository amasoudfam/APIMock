package handlers

import (
	"net/http"
	"on-air/models"
	"on-air/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Flight struct {
	DB *gorm.DB
}

type ListRequest struct {
	Origin      string `query:"org" validate:"required"`
	Destination string `query:"dest" validate:"required"`
	Date        string `query:"date" validate:"required,datetime=2006-01-02"`
}

type ListResponse struct {
	Flights []models.Flight `json:"flights"`
}

func (f *Flight) FlightsFromOrgToDestInDate(ctx echo.Context) error {
	var req ListRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid query parameters")
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	flights, err := repository.GetFlights(f.DB, req.Origin, req.Destination, req.Date)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	response := ListResponse{
		Flights: flights,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (f *Flight) Airplanes(ctx echo.Context) error {
	var res []string
	if err := f.DB.Model(&models.Flight{}).Select("airplane").Distinct().Find(&res).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (f *Flight) Cities(ctx echo.Context) error {
	var res []string
	if err := f.DB.Raw("? UNION ?",
		f.DB.Distinct("origin").Select("origin").Model(&models.Flight{}),
		f.DB.Distinct("destination").Select("destination").Model(&models.Flight{}),
	).Scan(&res).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

type ReserveRequest struct {
}

type ReserveResponse struct {
}

func (f *Flight) Reserve(ctx echo.Context) error {
	return nil
}
