package routes

import (
	"github.com/labstack/echo/v4"
	usercontroller "github.com/salamanderman234/simple-api/controllers/user.controller"
)

type userRoute struct {
	prefix string
}

func newUserRoute() *userRoute {
	return &userRoute{
		prefix: "/user",
	}
}

func (u *userRoute) getPrefix() string {
	return u.prefix
}

func (u *userRoute) register(router *echo.Echo) {
	router.GET(newUrl(u, ""), usercontroller.UserController)
	router.GET(newUrl(u, "/get"), usercontroller.UserGetController)
	router.GET(newUrl(u, "/new"), usercontroller.AddUserController)
}
