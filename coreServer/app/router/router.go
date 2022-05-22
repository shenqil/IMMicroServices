package router

import (
	"core-server/app/service"
	"core-server/pb"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

// RouterSet 注入Router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

// Router 路由管理器
type Router struct {
	DemoSrc    *service.Demo
	UserSrc    *service.User
	CaptchaSrc *service.Captcha
	TokenSrc   *service.Token
}

func (a *Router) Register(s grpc.ServiceRegistrar) {
	pb.RegisterDemoGreeterServer(s, a.DemoSrc)
	pb.RegisterUserGreeterServer(s, a.UserSrc)
	pb.RegisterCaptchaGreeterServer(s, a.CaptchaSrc)
	pb.RegisterTokenGreeterServer(s, a.TokenSrc)
}
