package main

import (
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/config"
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/route"
)

func main() {
	config.InitDB()
	e := route.SetupRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}
