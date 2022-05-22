package service

import (
	"context"
	"core-server/app/model/jwtauth"
	"core-server/pb"
	"errors"
	"github.com/google/wire"
)

// TokenSet 注入令牌
var TokenSet = wire.NewSet(wire.Struct(new(Token), "Auth"))

// Token 令牌
type Token struct {
	Auth *jwtauth.JWTAuth
	pb.UnimplementedTokenGreeterServer
}

func (t *Token) Generate(ctx context.Context, in *pb.TokenGenerateRequest) (*pb.TokenGenerateResult, error) {
	id := in.GetId()

	if id == "" {
		return nil, errors.New("ID不能为空")
	}

	token, err := t.Auth.GenerateToken(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.TokenGenerateResult{
		Token: &pb.Token{
			AccessToken: token.AccessToken,
			TokenType:   token.TokenType,
			ExpiresAt:   token.ExpiresAt,
		},
	}, nil
}

func (t Token) Destroy(ctx context.Context, in *pb.TokenDestroyRequest) (*pb.StatusResult, error) {
	accessToken := in.GetAccessToken()
	if accessToken == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("非法Token")
	}

	err := t.Auth.DestroyToken(ctx, accessToken)
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}

	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

func (t Token) Parse(ctx context.Context, in *pb.TokenParseRequest) (*pb.TokenParseResult, error) {
	accessToken := in.GetAccessToken()
	if accessToken == "" {
		return nil, errors.New("非法Token")
	}

	id, err := t.Auth.ParseID(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	return &pb.TokenParseResult{
		ID: id,
	}, nil
}
