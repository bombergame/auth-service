package main

import (
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/repositories/mysql"
	"github.com/bombergame/auth-service/services/grpc"
	"github.com/bombergame/auth-service/services/rest"
	"github.com/bombergame/common/auth/jwt"
	"github.com/bombergame/common/auth/randtoken"
	"github.com/bombergame/common/grpc"
	"github.com/bombergame/common/logs"
	"github.com/bombergame/common/rest"
	"github.com/bombergame/profiles-service/services/grpc"
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
		err := conn.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sessionRepository := mysql.NewSessionRepository(conn)

	profilesClient := profilesgrpc.NewClient(
		profilesgrpc.ClientConfig{
			ClientConfig: grpc.ClientConfig{
				ServiceHost: config.ProfilesServiceGrpcHost,
				ServicePort: config.ProfilesServiceGrpcPort,
			},
		},
		profilesgrpc.ClientComponents{
			ClientComponents: grpc.ClientComponents{
				Logger: logger,
			},
		},
	)

	if err := profilesClient.Connect(); err != nil {
		logger.Fatal(err)
		return
	}
	defer func() {
		err := profilesClient.Disconnect()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	restSrv := authrest.NewService(
		authrest.Config{
			Config: rest.Config{},
		},
		authrest.Components{
			Components: rest.Components{
				Logger: logger,
			},
			SessionRepository:   sessionRepository,
			AuthTokenManager:    jwt.NewTokenManager(config.TokenSignKey),
			RefreshTokenManager: randtoken.NewTokenManager(),
			ProfilesClient:      profilesClient,
		},
	)

	grpcSrv := authgrpc.NewService(
		authgrpc.ServiceConfig{
			ServiceConfig: grpc.ServiceConfig{},
		},
		authgrpc.ServiceComponents{
			ServiceComponents: grpc.ServiceComponents{
				Logger: logger,
			},
			SessionRepository: sessionRepository,
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
