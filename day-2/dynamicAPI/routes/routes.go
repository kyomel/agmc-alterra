package routes

import (
	"dynamicAPI/config"
	"dynamicAPI/controllers"
	"dynamicAPI/lib/database"

	"github.com/labstack/echo/v4"
)

var (
	Connection = config.InitDB()
	RepoUser   = database.Init(Connection)
)

func Init() *echo.Echo {
	e := echo.New()

	userController := controllers.Init(RepoUser)
	e.POST("/v1/users", userController.CreateUserControllers)
	e.GET("/v1/users", userController.GetUsersControllers)
	e.GET("/v1/users/:id", userController.GetUserByIDControllers)
	e.PUT("/v1/users/:id", userController.UpdateUserControllers)
	e.DELETE("/v1/users/:id", userController.DeleteUserControllers)

	return e
}
