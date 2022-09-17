package routes

import (
	"bookApp/config"
	"bookApp/controllers"
	"bookApp/helpers"

	"bookApp/lib/database"
	mid "bookApp/middleware"

	"github.com/labstack/echo/v4"
)

var (
	Connection = config.InitDB()
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	mid.LogMiddleware(e)

	// User Set
	databases := database.Init(Connection)
	controller := controllers.Init(databases)
	helper := helpers.Init(controller)

	// Book Set
	databasesBook := database.InitBook(Connection)
	controllerBook := controllers.InitBook(databasesBook)
	helperBook := helpers.InitBook(controllerBook)

	e.Validator = mid.NewCustomValidator()

	gJwt := e.Group("/jwt")
	mid.SetJwtMiddlewares(gJwt)

	e.POST("/v1/users", helper.CreateUser)
	e.POST("/v1/login", helper.UserLogin)
	e.GET("/v1/books/:id", helperBook.GetBookById)
	e.GET("/v1/books", helperBook.GetAllBooks)
	gJwt.PUT("/v1/users/:id", helper.UpdateUser)
	gJwt.GET("/v1/users/:id", helper.GetUserById)
	gJwt.GET("/v1/users", helper.GetAllUsers)
	gJwt.DELETE("/v1/users/:id", helper.DeleteUser)
	gJwt.PUT("/v1/books/:id", helperBook.UpdateBook)
	gJwt.DELETE("/v1/books/:id", helperBook.DeleteBook)
	gJwt.POST("/v1/books", helperBook.CreateBook)

	return e
}
