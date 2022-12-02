package routes

import (
	"github.com/labstack/echo/v4"
	controller "github.com/salamanderman234/simple-api/controllers/user"
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
	router.GET(newUrl(u, ""), controller.UserController)
	router.GET(newUrl(u, "/get"), controller.UserGetController)
}
