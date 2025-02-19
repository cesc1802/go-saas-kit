package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/goxiaoy/go-saas-kit/auth/jwt"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	http2 "github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/seed"
	"github.com/goxiaoy/uow"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/conf"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/data"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, seeder seed.Seeder) *kratos.App {
	if err := seeder.Seed(context.Background()); err != nil {
		panic(err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Services, bc.Data, logger, &jwt.TokenizerConfig{
		ExpireDuration: bc.Security.Jwt.ExpireIn.AsDuration(),
		Secret:         bc.Security.Jwt.Secret,
	}, &uow.Config{
		SupportNestedTransaction: false,
	}, uow2.NewGormConfig(bc.Data.Endpoints, data.ConnName),http2.NewDefaultWebMultiTenancyOption())
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
