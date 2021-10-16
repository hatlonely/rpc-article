package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hatlonely/go-kit/bind"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	_ "github.com/hatlonely/go-kit/logger/x"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/rpcx"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
	"github.com/hatlonely/rpc-article/internal/service"
)

var Version string

type Options struct {
	flag.Options

	Service     service.Options
	GrpcGateway rpcx.GrpcGatewayOptions

	Logger struct {
		Info logger.Options
		Grpc logger.Options
		Exec logger.Options
	}
}

func main() {
	var options Options
	refx.Must(flag.Struct(&options, refx.WithCamelName()))
	refx.Must(flag.Parse(flag.WithJsonVal()))
	if options.Help {
		fmt.Println(flag.Usage())
		return
	}
	if options.Version {
		fmt.Println(Version)
		return
	}

	if options.ConfigPath == "" {
		options.ConfigPath = "config/base.json"
	}

	cfg, err := config.NewConfigWithBaseFile(options.ConfigPath, refx.WithCamelName())
	refx.Must(err)
	refx.Must(cfg.Watch())
	defer cfg.Stop()

	refx.Must(bind.Bind(&options, []bind.Getter{
		flag.Instance(), bind.NewEnvGetter(bind.WithEnvPrefix("OPS")), cfg,
	}, refx.WithCamelName()))

	grpcLog, err := logger.NewLoggerWithOptions(&options.Logger.Grpc, refx.WithCamelName())
	refx.Must(err)
	infoLog, err := logger.NewLoggerWithOptions(&options.Logger.Info, refx.WithCamelName())
	refx.Must(err)
	execLog, err := logger.NewLoggerWithOptions(&options.Logger.Exec, refx.WithCamelName())
	refx.Must(err)
	infoLog.With("options", options).Info("init config success")

	_ = execLog

	svc, err := service.NewServiceWithOptions(&options.Service)
	refx.Must(err)

	grpcGateway, err := rpcx.NewGrpcGatewayWithOptions(&options.GrpcGateway)
	refx.Must(err)
	grpcGateway.SetLogger(infoLog, grpcLog)

	api.RegisterArticleServiceServer(grpcGateway.GRPCServer(), svc)
	refx.Must(grpcGateway.RegisterServiceHandlerFunc(api.RegisterArticleServiceHandlerFromEndpoint))
	grpcGateway.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	grpcGateway.Stop()
	infoLog.Info("server exit properly")
}