package main

import (
	"bookApp/routes"
)

func main() {
	e := routes.NewRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}
