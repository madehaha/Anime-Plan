package util

import "github.com/labstack/echo/v4"

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Success(c echo.Context, code int, data interface{}) error {
	return c.JSON(
		code, Response{
			Msg:  "",
			Data: data,
		},
	)
}

func Error(c echo.Context, code int, msg string) error {
	return c.JSON(
		code, Response{
			Msg:  msg,
			Data: nil,
		},
	)
}
