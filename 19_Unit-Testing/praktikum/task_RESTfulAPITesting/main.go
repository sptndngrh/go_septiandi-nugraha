package main

import (
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/config"
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	route.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
