package routes

import (
	"bookApp/database/config"
	"bookApp/internal/controllers"
	"bookApp/internal/helpers"

	mid "bookApp/internal/middleware"
	"bookApp/internal/repository"

	"github.com/labstack/echo/v4"
)

var (
	Connection      = config.InitDB()
	ConnectionMongo = config.ConnectDB()
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	mid.LogMiddleware(e)

	// User Set
	databases := repository.Init(Connection)
	controller := controllers.Init(databases)
	helper := helpers.Init(controller)

	// Book Set
	databasesBook := repository.InitBook(Connection)
	controllerBook := controllers.InitBook(databasesBook)
	helperBook := helpers.InitBook(controllerBook)

	// Review Set
	databasesReview := repository.InitReview(ConnectionMongo)
	controllerReview := controllers.InitReview(databasesReview)
	helperReview := helpers.InitReview(controllerReview)

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

	e.POST("/v1/review", helperReview.CreateReview)
	e.GET("/v1/reviews", helperReview.GetReviews)
	e.GET("/v1/review/:id", helperReview.GetReview)
	e.PUT("/v1/review/:id", helperReview.UpdateReview)
	e.DELETE("/v1/review/:id", helperReview.DeleteReview)

	return e
}
