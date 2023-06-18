package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func getSimonokuHandler(c echo.Context) error {
	id := c.Param("kaminoku_id")

	var list []Simonokudb
	if err := db.Select(&list, "SELECT simonoku.id as simonokuid ,kaminoku.id as kaminokuid,first,second,third,fourth,fifth,kaminoku.userid as kaminokuuser, simonoku.userid as simonokuuser FROM simonoku JOIN kaminoku ON kaminoku.id = simonoku.kaminokuid WHERE kaminokuid =?", id); errors.Is(err, sql.ErrNoRows) {
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

func postSimonokuHandler(c echo.Context) error {
	// 受け取りたい JSON を示す空の変数を先に用意する。
	data := &SimonokuReq{}
	// 受け取った JSON を data に代入する
	err := c.Bind(data)

	if err != nil { // エラーが発生した時、以下を実行
		// 正常でないためステータスコード 400 Bad Requestを返し、 エラーを文字列に変換して出力
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "なんか壊れた")
	}
	uu := u.String()
	un := c.Get("userName").(string)
	_, derr := db.Exec("INSERT INTO simonoku (id,fourth,fifth,kaminokuid,userid) VALUES (?, ?,?, ?,?)", uu, data.Content.Fourth, data.Content.Fifth, c.Param("kaminoku_id"), un)
	if derr != nil {
		log.Fatalf("failed to insert data: %s", err)
	}

	return c.NoContent(http.StatusCreated)
}
func getAllSimonokuHandler(c echo.Context) error {

	var list []Simonokudb
	if err := db.Select(&list, "SELECT simonoku.id as simonokuid ,kaminoku.id as kaminokuid,first,second,third,fourth,fifth,kaminoku.userid as kaminokuuser, simonoku.userid as simonokuuser FROM simonoku JOIN kaminoku ON kaminoku.id = simonoku.kaminokuid"); errors.Is(err, sql.ErrNoRows) {
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
