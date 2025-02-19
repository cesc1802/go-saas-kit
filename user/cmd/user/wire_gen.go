// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/goxiaoy/go-saas-kit/auth/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/user/internal_/biz"
	conf2 "github.com/goxiaoy/go-saas-kit/user/internal_/conf"
	"github.com/goxiaoy/go-saas-kit/user/internal_/data"
	"github.com/goxiaoy/go-saas-kit/user/internal_/seed"
	"github.com/goxiaoy/go-saas-kit/user/internal_/server"
	"github.com/goxiaoy/go-saas-kit/user/internal_/service"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/uow"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, userConf *conf2.UserConf, confData *conf2.Data, logger log.Logger, passwordValidatorConfig *biz.PasswordValidatorConfig, tokenizerConfig *jwt.TokenizerConfig, config *uow.Config, gormConfig *gorm.Config, webMultiTenancyOption *http.WebMultiTenancyOption) (*kratos.App, func(), error) {
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	dbOpener, cleanup := gorm.NewDbOpener()
	manager := uow2.NewUowManager(gormConfig, config, dbOpener)
	tenantStore := data.NewTenantStore()
	dbProvider := data.NewProvider(confData, gormConfig, dbOpener, manager, tenantStore, logger)
	dataData, cleanup2, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	passwordHasher := biz.NewPasswordHasher()
	userValidator := biz.NewUserValidator()
	passwordValidator := biz.NewPasswordValidator(passwordValidatorConfig)
	lookupNormalizer := biz.NewLookupNormalizer()
	userManager := biz.NewUserManager(userRepo, passwordHasher, userValidator, passwordValidator, lookupNormalizer, logger)
	userService := service.NewUserService(userManager)
	accountService := service.NewAccountService(userManager)
	roleRepo := data.NewRoleRepo(dataData)
	roleManager := biz.NewRoleManager(roleRepo, lookupNormalizer)
	authService := service.NewAuthService(userManager, roleManager, tokenizer, tokenizerConfig, passwordValidator)
	httpServer := server.NewHTTPServer(services, tokenizer, manager, webMultiTenancyOption, tenantStore, userService, accountService, authService, logger)
	grpcServer := server.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, userService, accountService, authService, logger)
	migrate := data.NewMigrate(dataData)
	roleSeed := biz.NewRoleSeed(roleManager)
	userSeed := biz.NewUserSeed(userManager, roleManager)
	fake := seed.NewFake(userManager)
	seeder := server.NewSeeder(userConf, manager, migrate, roleSeed, userSeed, fake)
	app := newApp(logger, httpServer, grpcServer, seeder)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
