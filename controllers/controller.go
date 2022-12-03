package controller

import (
	"github.com/labstack/echo"
)

type Controller interface {
	Add(echo.Context) error
	Read(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}
