package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getUserKaminokuHandler(c echo.Context) error {
	id := c.Get("userName").(string)
	fmt.Println("test")

	var kami Kaminokudb
	if err := db.Get(&kami, "SELECT * FROM kaminoku WHERE userid = ?", id); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "error...")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	res := Kaminoku{Id: kami.Id, Content: Fsf{kami.First, kami.Second, kami.Third}, Userid: kami.Userid}

	return c.JSON(http.StatusOK, res)
}
func getUserSimonokuHandler(c echo.Context) error {
	id := c.Get("userName").(string)

	var list []Simonokudb
	if err := db.Select(&list, "SELECT simonoku.id as simonokuid ,kaminoku.id as kaminokuid,first,second,third,fourth,fifth,kaminoku.userid as kaminokuuser, simonoku.userid as simonokuuser FROM simonoku JOIN kaminoku ON kaminoku.id = simonoku.kaminokuid WHERE simonoku.userid = ?", id); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "error...")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}
	var res []TankaRes
	for _, k := range list {
		res = append(res, TankaRes{
			Kaminoku: Kaminoku{
				Id:      k.Kaminokuid,
				Content: Fsf{k.First, k.Second, k.Third},
				Userid:  k.Kaminokuuserid,
			},
			Simonoku: Simonoku{
				Id:      k.Simonokuid,
				Content: Ss{k.Fourth, k.Fifth},
				Userid:  k.Simonokuuserid,
			},
		})
	}

	return c.JSON(http.StatusOK, res)
}
