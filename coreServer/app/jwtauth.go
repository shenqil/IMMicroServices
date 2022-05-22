package app

import (
	"core-server/app/model/jwtauth"
	"core-server/app/model/redisx"
	"core-server/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// InitAuth 初始化用户认证
func InitAuth() (*jwtauth.JWTAuth, func(), error) {

	cfg := config.C.JWTAuth
	defaultOptions := &jwtauth.Options{
		TokenType:     "Bearer",
		Expired:       cfg.Expired,
		SigningMethod: jwt.SigningMethodHS512,
		SigningKey:    []byte(cfg.SigningKey),
		Keyfunc: func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}
			return []byte(cfg.SigningKey), nil
		},
	}

	switch cfg.SigningMethod {
	case "HS256":
		defaultOptions.SigningMethod = jwt.SigningMethodHS256
	case "HS384":
		defaultOptions.SigningMethod = jwt.SigningMethodHS384
	default:
		defaultOptions.SigningMethod = jwt.SigningMethodHS512
	}

	rcfg := config.C.Redis
	store := redisx.NewStore(&redisx.Config{
		Addr:      rcfg.Addr,
		Password:  rcfg.Password,
		DB:        cfg.RedisDB,
		KeyPrefix: cfg.RedisPrefix,
	})

	auth := &jwtauth.JWTAuth{
		Opts:  defaultOptions,
		Store: store,
	}

	cleanFunc := func() {
		auth.Release()
	}

	return auth, cleanFunc, nil
}
