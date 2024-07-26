package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartWeb() {
	e := echo.New()

	e.GET("/api/incidents", func(c echo.Context) error {
		cacheData, err := Cache.Get(c.Request().Context(), "peto-data")

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch data")
		}

		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")

		return c.String(http.StatusOK, cacheData)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
