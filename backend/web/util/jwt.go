package util

import (
	"backend/internal/config"
	"backend/internal/logger"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtUtil struct {
	cfg config.AppConfig
}

type JwtUserClaims struct {
	Uid uint32 `json:"uid"`
	Gid uint8  `json:"gid"`
	jwt.StandardClaims
}

func NewJwtUtil(appConfig config.AppConfig) JwtUtil {
	logger.Debug("Here")
	return JwtUtil{cfg: appConfig}
}

func (j JwtUtil) GenerateNewToken(uid uint32, gid uint8) string {
	claims := &JwtUserClaims{
		Uid: uid,
		Gid: gid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	logger.Info("Here")
	token, _ := t.SignedString([]byte(j.cfg.JwtSecret))
	return token
}

func (j JwtUtil) ParseToken(token string) (*JwtUserClaims, error) {
	parseWithClaims, err := jwt.ParseWithClaims(token, &JwtUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.cfg.JwtSecret), nil
	})
	if parseWithClaims != nil {
		if claims, ok := parseWithClaims.Claims.(*JwtUserClaims); ok && parseWithClaims.Valid {
			return claims, err
		}
	}
	return nil, err
}
