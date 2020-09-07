package main

import (
	"hology/database"
	"hology/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1200"))
}
