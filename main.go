package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var (
	db *sqlx.DB
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	conf := mysql.Config{
		User:      os.Getenv("DB_USERNAME"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_DATABASE"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conntected")
	db = _db

	e := echo.New()

	e.POST("/login", getLoginHandler)
	e.POST("/register", getRegisterHandler)

	e.GET("/kaminoku", getKaminokuHandler)
	e.POST("/kaminoku", postKaminokuHandler)
	e.GET("/kaminoku/:kaminoku_id", getKaminokuDetailHandler)

	e.GET("/kaminoku/:kaminoku_id/simonoku", getSimonokuHandler)
	e.POST("/kaminoku/:kaminoku_id/simonoku", postSimonokuHandler)
	e.GET("/simonoku", getAllSimonokuHandler)

	e.GET("/user/:user_id/kaminoku", getUserKaminokuHandler)
	e.GET("/user/:user_id/simonoku", getUserSimonokuHandler)
	e.Start(":8080")
}
