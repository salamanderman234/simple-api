package routes

import "github.com/labstack/echo/v4"

type route interface {
	register(*echo.Echo)
	getPrefix() string
}

func Init() *echo.Echo {
	routes := []route{
		newUserRoute(),
	}
	router := echo.New()

	for _, route := range routes {
		route.register(router)
	}

	return router
}

func newUrl(model route, path string) string {
	return model.getPrefix() + path
}
