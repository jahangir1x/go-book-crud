package routes

import (
	"book-crud/pkg/controllers"
	"github.com/labstack/echo/v4"
)

type AuthRoutes struct {
	echo    *echo.Echo
	authCtr controllers.AuthController
}

func NewAuthRoutes(echo *echo.Echo, authCtr controllers.AuthController) *AuthRoutes {
	return &AuthRoutes{
		echo:    echo,
		authCtr: authCtr,
	}
}

func (ac *AuthRoutes) InitAuthRoutes() {
	e := ac.echo
	ac.initAuthRoutes(e)
}

func (ac *AuthRoutes) initAuthRoutes(e *echo.Echo) {

	e.POST("/login", ac.authCtr.Login)
	e.POST("/signup", ac.authCtr.Signup)
}
