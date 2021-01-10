package main

import (
	"gintemplate/app/controllers"
	"gintemplate/app/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	controllers.RooterInit()
	models.DbInit()
}
