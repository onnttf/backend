package main

import (
	"backend/util/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type gasoline struct {
	Price_92 string `json:"price_92"`
	Price_95 string `json:"price_95"`
	Price_98 string `json:"price_98"`
}

type City struct {
	Zhcity string `json:"zhcity"`
	City   string `json:"city"`
}

func main() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:    true,
		LogURI:       true,
		LogRequestID: true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogStatus:    true,
		LogError:     true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l := logger.Get()
			if v.Error == nil {
				l.Info().
					Str("remote_ip", v.RemoteIP).
					Str("method", v.Method).
					Str("uri", v.URI).
					Int("status", v.Status).
					Dur("latency", v.Latency).
					Str("request_id", v.RequestID).
					Send()
			} else {
				l.Error().
					Str("remote_ip", v.RemoteIP).
					Str("method", v.Method).
					Str("uri", v.URI).
					Int("status", v.Status).
					Dur("latency", v.Latency).
					Str("request_id", v.RequestID).
					Err(v.Error).
					Send()
			}
			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
