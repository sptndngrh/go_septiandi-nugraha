package main

import (
	"go_septiandi-nugraha/18_Middleware/praktikum/task/config"
	"go_septiandi-nugraha/18_Middleware/praktikum/task/route"
)

func main() {
	config.InitDB()
	e := route.SetupRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}
