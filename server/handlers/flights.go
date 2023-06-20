package handlers

import (
	"net/http"
	"on-air/models"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Flight struct {
	DB *gorm.DB
}

type ListRequest struct {
	Origin      string    `json:"origin" validate:"required"`
	Destination string    `json:"destination" validate:"required"`
	Date        time.Time `json:"date" query:"date" validate:"required,datetime=2006-01-02"`
}

type ListResponse struct {
	Flights []models.Flight `json:"flights"`
}

func (f *Flight) List(ctx echo.Context) error {
	var req ListRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// TODO Fetch flights from database and send with response

	response := ListResponse{
		Flights: []models.Flight{},
	}

	return ctx.JSON(http.StatusOK, response)
}
