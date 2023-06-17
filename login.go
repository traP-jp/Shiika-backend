package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func postLoginHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	var test User
	if err := db.Get(&test, "SELECT * FROM user WHERE name='ramdos'"); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No such city Name = %s", cityName))
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, test.Password)
}
func postRegisterHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	var test User
	if err := db.Get(&test, "SELECT * FROM user WHERE name='ramdos'"); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No such city Name = %s", cityName))
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, test.Password)
}
