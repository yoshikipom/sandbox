package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Entry struct {
	ID int `json:"id"`
}

func main() {
	e := echo.New()

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mydb")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	e.GET("/data", func(c echo.Context) error {
		rows, err := db.Query("SELECT id FROM t")
		if err != nil {
			return err
		}
		defer rows.Close()

		entries := []Entry{}
		for rows.Next() {
			var ent Entry
			if err := rows.Scan(&ent.ID); err != nil {
				return err
			}
			entries = append(entries, ent)
		}

		return c.JSON(http.StatusOK, entries)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
