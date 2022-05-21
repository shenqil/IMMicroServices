package app

import (
	"context"
	"core-server/app/router"
	"core-server/config"
	"core-server/util/grpcLogrus"
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

	InitGRPCServer(injector.Router)

	return nil
}

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
