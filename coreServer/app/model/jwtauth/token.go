package jwtauth

import (
	"context"
	"core-server/app/model/redisx"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenInfo struct {
	//	访问令牌
	AccessToken string
	//	令牌类型
	TokenType string
	// 令牌到期时间
	ExpiresAt int64
}

type Options struct {
	TokenType     string
	SigningMethod jwt.SigningMethod
	SigningKey    interface{}
	Keyfunc       jwt.Keyfunc
	Expired       int
}

// JWTAuth jwt认证
type JWTAuth struct {
	Opts  *Options
	Store *redisx.Store
}

// GenerateToken 生成令牌
func (a *JWTAuth) GenerateToken(ctx context.Context, id string) (*TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.Opts.Expired) * time.Second).Unix()

	token := jwt.NewWithClaims(a.Opts.SigningMethod, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   id,
	})

	tokenString, err := token.SignedString(a.Opts.SigningKey)
	if err != nil {
		return nil, err
	}

	tokenInfo := &TokenInfo{
		ExpiresAt:   expiresAt,
		TokenType:   a.Opts.TokenType,
		AccessToken: tokenString,
	}

	return tokenInfo, nil
}

// 解析令牌
func (a *JWTAuth) parseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, a.Opts.Keyfunc)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(*jwt.StandardClaims), nil
}

// DestroyToken 销毁令牌
func (a *JWTAuth) DestroyToken(ctx context.Context, tokenString string) error {
	claims, err := a.parseToken(tokenString)
	if err != nil {
		return err
	}

	//	将未过期的令牌放入,表示该token已被销毁
	expired := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())
	return a.Store.Set(ctx, tokenString, expired)
}

// ParseID 解析ID
func (a *JWTAuth) ParseID(ctx context.Context, tokenString string) (string, error) {
	if tokenString == "" {
		return "", errors.New("invalid token")
	}

	claims, err := a.parseToken(tokenString)
	if err != nil {
		return "", err
	}

	// 判断此token是否被销毁
	if exists, err := a.Store.Check(ctx, tokenString); err != nil {
		return "", err
	} else if exists {
		return "", errors.New("invalid token")
	}

	return claims.Subject, nil
}

// Release 释放资源
func (a *JWTAuth) Release() error {
	return a.Store.Close()
}
