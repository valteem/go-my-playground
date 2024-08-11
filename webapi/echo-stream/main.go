package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type WeatherSpot struct {
	Timestamp string
	Temp      float32
	Windspeed float32
	Pressure  float32
}

var (
	spots = []WeatherSpot{
		{"20240811170800", 19.1, 4.8, 741.1},
		{"20240811171230", 18.8, 4.7, 740.5},
		{"20240811172445", 18.7, 4.5, 740.1},
		{"20240811173115", 18.5, 4.1, 740.0},
		{"20240811173535", 18.4, 4.2, 739.8}}
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Content-Type: application/json
		c.Response().WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(c.Response())
		for _, spot := range spots {
			if err := encoder.Encode(spot); err != nil {
				return err
			}
			c.Response().Flush() // flush buffered data to the client (http.ResponseController.Flush())
			time.Sleep(time.Second * 2)
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":44568"))
}
