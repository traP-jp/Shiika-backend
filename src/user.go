package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getUserKaminokuHandler(c echo.Context) error {
	id := c.Get("userName").(string)

	var kaminoku []Kaminoku
	if err := db.Select(&kaminoku, "SELECT * FROM kaminoku WHERE userid =?", id); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "error...")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, kaminoku)
}
func getUserSimonokuHandler(c echo.Context) error {
	id := c.Get("userName").(string)

	var list []TankaRes
	if err := db.Select(&list, "SELECT simonoku.id,name as simonoku,content AS kaminoku,kaminoku.userid AS kaminokuuser,simonoku.userid AS simonokuuser FROM simonoku JOIN kaminoku ON kaminoku.id = simonoku.kaminokuid WHERE simonoku.userid =?", id); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "error...")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, list)
}
