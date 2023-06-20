package handlers

import (
	"net/http"
	"on-air/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Flight struct {
	DB *gorm.DB
}

type ListRequest struct {
	Origin      string `json:"origin" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Date        string `json:"date" query:"date" validate:"required,datetime=2006-01-02"`
}

type ListResponse struct {
	Flights []models.Flight `json:"flights"`
}

func (f *Flight) List(ctx echo.Context) error {
	var req ListRequest

	// FIXME: c.Bind(&req) does not work
	req.Origin = ctx.QueryParam("origin")
	req.Destination = ctx.QueryParam("destination")
	req.Date = ctx.QueryParam("date")

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var flights []models.Flight
	if err := f.DB.Where("origin = ? AND destination = ? AND DATE(started_at) = ?", req.Origin, req.Destination, req.Date).Find(&flights).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	response := ListResponse{
		Flights: flights,
	}

	return ctx.JSON(http.StatusOK, response)
}
