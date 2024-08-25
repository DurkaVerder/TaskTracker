package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func Run(e *echo.Echo, db *sql.DB) {

	e.GET("/tasks", func(c echo.Context) error {

		return nil
	})

	e.POST("/addTask", func(c echo.Context) error {

		return nil
	})

	e.PUT("/changeTask", func(c echo.Context) error {

		return nil
	})

	e.DELETE("/deleteTask", func(c echo.Context) error {

		return nil
	})

	e.Start(":2222")

}
