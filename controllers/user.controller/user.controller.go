package usercontroller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/salamanderman234/simple-api/models"
)

func UserController(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Berhasil mendaftarkan user route !")
}
func UserGetController(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "user 1 user 2 user 3")
}

func AddUserController(ctx echo.Context) error {
	newUser := model.User{
		Name:    "saya 1",
		Address: "Jalan Pratu Made Rampug",
		Age:     12,
	}
	err := newUser.CreateMe()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "tidak sukses menambahkan user")
	}
	return ctx.String(http.StatusOK, "berhasil menambahkan user 1")
}
