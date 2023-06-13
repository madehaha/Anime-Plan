package response

import "github.com/labstack/echo/v4"

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Success(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(code, Response{
		Msg:  "",
		Data: data,
	})
}

func Error(ctx echo.Context, code int, msg string) error {
	return ctx.JSON(code, Response{
		Msg:  msg,
		Data: nil,
	})
}
