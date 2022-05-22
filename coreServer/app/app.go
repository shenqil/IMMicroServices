package app

import (
	"context"
	"core-server/app/router"
	"core-server/config"
	"core-server/util/grpcLogrus"
	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

// Run 运行服务
func Run(ctx context.Context) error {

	// 初始化依赖注入器
	injector, injectorCleanFunc, err := BuildInjector()
	if err != nil {
		injectorCleanFunc()
		logrus.Fatalf("[app][BuildInjector] err :%v", err)
		return err
	}

	InitCaptcha()
	InitGRPCServer(injector.Router)

	return nil
}

// InitGRPCServer 初始化GRPC服务
func InitGRPCServer(r *router.Router) {
	grpclog.SetLoggerV2(grpcLogrus.NewLogrus(logrus.New()))
	// 监听端口
	lis, err := net.Listen("tcp", config.C.GRPC.Addr())
	if err != nil {
		logrus.Errorf("[grpc][Listen] err:%v", err)
		return
	}
	logrus.Printf("[main][GRPC]服务器监听 %v", lis.Addr())
	// 开启GRPC服务，并注册对应路由
	grpcServer := grpc.NewServer()
	r.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Errorf("[grpc][Serve] err:%v", err)
		return
	}
}

// InitCaptcha 初始化图形验证码
func InitCaptcha() {
	cfg := config.C.Captcha
	if cfg.Store == "redis" {
		rc := config.C.Redis
		captcha.SetCustomStore(store.NewRedisStore(&redis.Options{
			Addr:     rc.Addr,
			Password: rc.Password,
			DB:       cfg.RedisDB,
		}, captcha.Expiration, logrus.StandardLogger(), cfg.RedisPrefix))
	}
}
