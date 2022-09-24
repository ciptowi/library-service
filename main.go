package main

import (
	"flag"
	"library-sevice/database/migration"
	"library-sevice/internal/factory"
	"library-sevice/internal/http"
	"os"

	"github.com/joho/godotenv"

	"library-sevice/database"

	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
func main() {

	database.CreateConnection()
	migration.Migrate()

	var m string // for check migration

	flag.StringVar(
		&m,
		"migrate",
		"run",
		`this argument for check if user want to migrate table, rollback table, or status migration

to use this flag:
	use -migrate=migrate for migrate table
	use -migrate=rollback for rollback table
	use -migrate=status for get status migration`,
	)
	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
		return
	} else if m == "rollback" {
		migration.Rollback()
		return
	} else if m == "status" {
		migration.Status()
		return
	}

	e := echo.New()

	f := factory.NewFactory()
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
