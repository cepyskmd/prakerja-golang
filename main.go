package main

import (
	"myapp/configs"
	"myapp/routes"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	configs.ConnectDatabase()
	e := routes.Init()
	e.Start(":8000")
}
