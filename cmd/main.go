package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maantje/postcode/handler"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	db, err := sql.Open("sqlite3", "file:database/addresses.sqlite?cache=shared")

	if err != nil {
		panic("Failed to connect to database")
	}

	db.SetMaxOpenConns(1)

	postcodeHandler := &handler.PostcodeHandler{
		Db: db,
	}

	docHandler := &handler.DocHandler{}

	e.GET("/postcode/:postcode/:houseNumber", postcodeHandler.Search)
	e.GET("/docs", docHandler.Docs)

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":80"))
}
