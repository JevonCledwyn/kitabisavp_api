package main

import (
	"kitabisavp/db"
	"kitabisavp/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":7070"))
}
