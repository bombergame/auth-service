package main

import (
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/repositories/mysql"
	"github.com/bombergame/auth-service/services/grpc"
	"github.com/bombergame/auth-service/services/rest"
	"github.com/bombergame/common/auth/jwt"
	"github.com/bombergame/common/auth/randtoken"
	grpcservice "github.com/bombergame/common/grpc"
	"github.com/bombergame/common/logs"
	restservice "github.com/bombergame/common/rest"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	conn := mysql.NewConnection()
	if err := conn.Open(); err != nil {
		logger.Fatal(err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Error(err)
		}
	}()

	restSrv := rest.NewService(
		rest.Config{
			Config: restservice.Config{},
		},
		rest.Components{
			Components: restservice.Components{
				Logger: logger,
			},
			AuthTokenManager:    jwt.NewTokenManager(config.TokenSignKey),
			RefreshTokenManager: randtoken.NewTokenManager(),
			SessionRepository:   mysql.NewSessionRepository(conn),
		},
	)

	grpcSrv := grpc.NewService(
		grpc.Config{
			Config: grpcservice.Config{},
		},
		grpc.Components{
			Components: grpcservice.Components{
				Logger: logger,
			},
			SessionRepository: mysql.NewSessionRepository(conn),
		},
	)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		if err := restSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	go func() {
		if err := grpcSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-ch

	if err := restSrv.Shutdown(); err != nil {
		logger.Fatal(err)
	}

	if err := grpcSrv.Shutdown(); err != nil {
		logger.Fatal(err)
	}
}
