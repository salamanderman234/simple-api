package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserController(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Berhasil mendaftarkan user route !")
}
func UserGetController(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "user 1 user 2 user 3")
}
