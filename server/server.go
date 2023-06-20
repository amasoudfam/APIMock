package server

import (
	"fmt"
	"net/http"
	"on-air/config"
	"on-air/server/handlers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func SetupServer(cfg *config.Config, db *gorm.DB, port string) error {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	Flight := &handlers.Flight{
		DB: db,
	}

	e.GET("/flights", Flight.List)

	return e.Start(fmt.Sprintf(":%s", port))
}
