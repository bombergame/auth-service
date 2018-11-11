package main

import (
	profilesgrpc "github.com/bombergame/auth-service/clients/profiles-service/grpc"
	"github.com/bombergame/auth-service/services/rest"
	"github.com/bombergame/auth-service/utils/jwt"
	"github.com/bombergame/common/logs"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	profilesGrpc := profilesgrpc.NewClient(
		&profilesgrpc.Config{
			Logger: logger,
		},
	)

	defer profilesGrpc.Disconnect()
	if err := profilesGrpc.Connect(); err != nil {
		logger.Fatal(err)
		return
	}

	restSrv := rest.NewService(
		&rest.Config{
			Logger:       logger,
			TokenManager: jwt.NewTokenManager(),
			ProfilesGrpc: profilesGrpc,
		},
	)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		if err := restSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-ch

	if err := restSrv.Shutdown(); err != nil {
		logger.Fatal(err)
	}
}
