package routes

import "github.com/labstack/echo/v4"

type route interface {
	register(*echo.Echo)
	getPrefix() string
}

func Init() *echo.Echo {
	routesList := []route{
		newUserRoute(),
	}
	router := echo.New()

	for _, element := range routesList {
		element.register(router)
	}

	return router
}

func newUrl(model route, path string) string {
	return model.getPrefix() + path
}
