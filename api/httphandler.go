package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "hello linda！")
}
