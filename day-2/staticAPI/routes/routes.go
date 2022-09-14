package routes

import (
	"staticAPI/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// book routes
	e.GET("/v1/books", controllers.GetBookAllControllers)
	e.GET("/v1/books/:id", controllers.GetBooksByID)
	e.POST("/v1/books", controllers.CreateBookControllers)
	e.PUT("/v1/books/:id", controllers.UpdateBooksByID)
	e.DELETE("/v1/books/:id", controllers.DeleteBooks)

	return e
}
