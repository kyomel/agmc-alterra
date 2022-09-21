package main

import (
	"bookApp/internal/routes"
)

func main() {
	e := routes.NewRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}
