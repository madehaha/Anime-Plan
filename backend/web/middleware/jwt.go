package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/internal/logger"
	"backend/web/util"
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

func (jm JwtMiddleware) extractAndCheckToken(ctx echo.Context) (
	token string, claims *util.JwtUserClaims, err error,
) {
	logger.Info("Here")
	token = ctx.Request().Header.Get("Authorization")
	fmt.Println(token)
	bearerLength := len(PREFIX)
	if len(token) < bearerLength {
		err = errors.New("please provide token")
		return
	} else {
		token = token[bearerLength:]
		claims, err = jm.Util.ParseToken(token)
		if err != nil {
			logger.Error("JWT parse token failed")
			return
		}
	}
	return
}

func (jm JwtMiddleware) UserJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, claims, err := jm.extractAndCheckToken(c)
		if err != nil {
			logger.Error("extract token or check token error")
			return util.Error(c, http.StatusUnauthorized, err.Error())
		}
		c.Set("uid", claims.Uid)
		c.Set("gid", claims.Gid)
		return next(c)
	}
}

func (jm JwtMiddleware) WikiJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, claims, err := jm.extractAndCheckToken(c)
		if err != nil {
			logger.Error("extract token or check token error")
			return util.Error(c, http.StatusUnauthorized, err.Error())
		}
		if claims.Gid == 0 {
			logger.Error("Permission denied")
			return util.Error(c, http.StatusUnauthorized, "Permission denied. You need to upgrade your user group")
		}
		c.Set("uid", claims.Uid)
		c.Set("gid", claims.Gid)
		return next(c)
	}
}
