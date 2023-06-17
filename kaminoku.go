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

func getKaminokuHandler(c echo.Context) error {

	var list []Kaminoku
	if err := db.Select(&list, "SELECT * FROM kaminoku"); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "error")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, list)
}

func postKaminokuHandler(c echo.Context) error {
	// 受け取りたい JSON を示す空の変数を先に用意する。
	data := &KaminokuReq{}
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

	_, derr := db.Exec("INSERT INTO kaminoku (id,content,userid) VALUES (?, ?, ?)", uu, data.Content, c.Get("userName").(string))
	if derr != nil {
		log.Fatalf("failed to insert city data: %s", err)
	}

	return c.NoContent(http.StatusCreated)
}
func getKaminokuDetailHandler(c echo.Context) error {
	id := c.Param("kaminoku_id")
	fmt.Println("test")

	var kaminoku Kaminoku
	if err := db.Get(&kaminoku, "SELECT * FROM kaminoku WHERE id =?", id); errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "error...")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, kaminoku)
}
