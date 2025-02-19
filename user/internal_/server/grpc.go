package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-saas-kit/auth/jwt"
	"github.com/goxiaoy/go-saas-kit/auth/middleware/authentication"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/server"
	"github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/user/api"
	v13 "github.com/goxiaoy/go-saas-kit/user/api/account/v1"
	v14 "github.com/goxiaoy/go-saas-kit/user/api/auth/v1"
	v12 "github.com/goxiaoy/go-saas-kit/user/api/user/v1"
	"github.com/goxiaoy/go-saas-kit/user/internal_/service"
	"github.com/goxiaoy/go-saas/common"
	http2 "github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/kratos/middleware"
	uow2 "github.com/goxiaoy/uow"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Services, tokenizer jwt.Tokenizer, ts common.TenantStore, uowMgr uow2.Manager,mOpt *http2.WebMultiTenancyOption, user *service.UserService, account *service.AccountService, auth *service.AuthService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			authentication.ServerExtractAndAuth(logger, tokenizer),
			uow.Uow(logger, uowMgr),
			middleware.Server(mOpt, nil, ts),
		),
	}
	opts = server.PatchGrpcOpts(opts, api.ServiceName, c)
	srv := grpc.NewServer(opts...)
	v12.RegisterUserServiceServer(srv, user)
	v13.RegisterAccountServer(srv, account)
	v14.RegisterAuthServer(srv, auth)
	return srv
}
