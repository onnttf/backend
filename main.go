package main

import (
	"backend/controller/petrol_price"
	"backend/dal"
	"backend/util/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	l := logger.Get()
	err := dal.InitMySQL()
	if err != nil {
		l.Info().Msg("222")
		l.Error().Err(err).Msg("init MySQL failed")
		panic(err)
	}
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
	e.Use(middleware.CORS())
	e.GET("/api/petrol_price", petrol_price.Query)
	if err := e.Start(":1323"); err != nil {
		l.Error().Err(err).Msg("start echo server failed")
		panic(err)
	}
}
