package middleware

import (
	"backend/internal/logger"
	"backend/web/util"
	"github.com/labstack/echo/v4"
)

const PREFIX = "Bearer "

type JwtMiddleware struct {
	Util util.JwtUtil
}

func NewJwtMiddleware(jwtUtil util.JwtUtil) JwtMiddleware {
	return JwtMiddleware{
		Util: jwtUtil,
	}
}

func (jm JwtMiddleware) extractAndCheckToken(ctx echo.Context) (token string, claims *util.JwtUserClaims, err error) {
	token = ctx.Request().Header.Get("Authorization")
	bearerLength := len(PREFIX)
	if len(token) < bearerLength {
		// TODO Error handling
		return
	}
	token = token[bearerLength:]
	claims, err = jm.Util.ParseToken(token)
	if err != nil {
		logger.Error("JWT parse token failed")
		return
	}
	return
}

func (jm JwtMiddleware) UserJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, _, err := jm.extractAndCheckToken(ctx)
		if err != nil {
			logger.Error("extract token or check token error")
			return err
		}
		return nil
	}
}
