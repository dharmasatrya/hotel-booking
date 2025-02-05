package routes

import (
	"github.com/dharmasatrya/hotel-booking/gateway-service/config"
	"github.com/dharmasatrya/hotel-booking/gateway-service/src/controller"
	"github.com/dharmasatrya/hotel-booking/gateway-service/src/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	userClient, err := config.InitUserServiceClient()
	if err != nil {
		e.Logger.Fatalf("did not connect: %v", err)
	}

	userService := service.NewUserService(userClient)
	userController := controller.NewUserController(userService)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	user := e.Group("/users")
	user.POST("/register/supporters", userController.RegisterUser)

	return e
}
