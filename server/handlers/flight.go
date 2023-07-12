package handlers

import (
	"net/http"
	"on-air/repository"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Flight struct {
	DB *gorm.DB
}

type GetFlightsRequest struct {
	Origin      string `query:"origin" validate:"required"`
	Destination string `query:"destination" validate:"required"`
	Date        string `query:"date" validate:"required,datetime=2006-01-02"`
}

type FlightFields struct {
	Number        string    `json:"number"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	Airplane      string    `json:"airplane"`
	Airline       string    `json:"airline"`
	Capacity      int       `json:"capacity"`
	EmptyCapacity int       `json:"empty_capacity"`
	Price         int       `json:"price"`
	StartedAt     time.Time `json:"started_at"`
	FinishedAt    time.Time `json:"finished_at"`
	Penalties     datatypes.JSON
}

func (f *Flight) GetFlights(ctx echo.Context) error {
	var req GetFlightsRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid query parameters")
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, "Input data is not valid")
	}

	flights, err := repository.GetFlights(f.DB, req.Origin, req.Destination, req.Date)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get flights. Please try again later.")
	}

	response := make([]FlightFields, len(flights))
	for i, flight := range flights {
		response[i] = FlightFields{
			Number:        flight.Number,
			Origin:        flight.Origin,
			Destination:   flight.Destination,
			Airplane:      flight.Airplane,
			Airline:       flight.Airline,
			Capacity:      flight.Capacity,
			EmptyCapacity: flight.EmptyCapacity,
			Price:         flight.Price,
			StartedAt:     flight.StartedAt,
			FinishedAt:    flight.FinishedAt,
			Penalties:     flight.Penalties,
		}
	}

	return ctx.JSON(http.StatusOK, response)
}

type AirplanesResponse struct {
	Airplanes []string
}

func (f *Flight) Airplanes(ctx echo.Context) error {
	airplanes, err := repository.GetAirplanes(f.DB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get airplanes. Please try again later.")
	}

	return ctx.JSON(http.StatusOK, AirplanesResponse{
		Airplanes: airplanes,
	})
}

func (f *Flight) Cities(ctx echo.Context) error {
	cities, err := repository.GetCities(f.DB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get cities. Please try again later.")
	}

	return ctx.JSON(http.StatusOK, cities)
}

func (f *Flight) Dates(ctx echo.Context) error {
	dates, err := repository.GetDates(f.DB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get dates. Please try again later.")
	}

	return ctx.JSON(http.StatusOK, dates)
}

func (f *Flight) Flight(ctx echo.Context) error {
	number := ctx.Param("number")
	flight, err := repository.GetFlight(f.DB, number)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get the Flight. Please try again later.")
	}

	return ctx.JSON(http.StatusOK, FlightFields{
		Number:        flight.Number,
		Origin:        flight.Origin,
		Destination:   flight.Destination,
		Airplane:      flight.Airplane,
		Airline:       flight.Airline,
		EmptyCapacity: flight.EmptyCapacity,
		Price:         flight.Price,
		StartedAt:     flight.StartedAt,
		FinishedAt:    flight.FinishedAt,
		Penalties:     flight.Penalties,
	})
}

type ReserveRequest struct {
	Number string `json:"number" validate:"required"`
	Count  int    `json:"count" validate:"required"`
}

type ReserveResponse struct {
	Status  bool
	Message string
}

func (f *Flight) Reserve(ctx echo.Context) error {
	var req ReserveRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request parameters")
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, "input data is not valid")
	}

	status, err := repository.DecrementEmptyCapacity(f.DB, req.Number, req.Count)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to reserve the flight. Please try again later.")
	}

	res := ReserveResponse{
		Status:  status,
		Message: "",
	}

	if !status {
		res.Message = "Flight reservation failed. No available capacity."
		return ctx.JSON(http.StatusExpectationFailed, res)
	}

	res.Message = "Flight reservation was successful."
	return ctx.JSON(http.StatusOK, res)
}

type RefundRequest struct {
	Number string `json:"number" validate:"required"`
	Count  int    `json:"count" validate:"required"`
}

type RefundResponse struct {
	Status  bool
	Message string
}

func (f *Flight) Refund(ctx echo.Context) error {
	var req ReserveRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request parameters")
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, "input data is not valid")
	}

	status, err := repository.IncrementEmptyCapacity(f.DB, req.Number, req.Count)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to refund the flight. Please try again later.")
	}

	res := ReserveResponse{
		Status:  status,
		Message: "",
	}

	if !status {
		res.Message = "Flight refund failed."
		return ctx.JSON(http.StatusExpectationFailed, res)
	}

	res.Message = "Flight refund was successful."
	return ctx.JSON(http.StatusOK, res)
}
