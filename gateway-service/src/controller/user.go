package controller

import (
	"net/http"

	"github.com/dharmasatrya/hotel-booking/gateway-service/src/service"

	"github.com/dharmasatrya/hotel-booking/user-service/dto"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (us *userController) RegisterUser(c echo.Context) error {
	var payload dto.RegisterRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err1 := us.userService.RegisterUser(payload)

	if err1 != nil {
		return c.JSON(http.StatusBadRequest, err1)
	}

	return c.JSON(http.StatusCreated, res)
}

func (us *userController) LoginUser(c echo.Context) error {
	var payload dto.LoginRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err1 := us.userService.LoginUser(payload)

	if err1 != nil {
		return c.JSON(http.StatusBadRequest, err1)
	}

	return c.JSON(http.StatusOK, res)
}
