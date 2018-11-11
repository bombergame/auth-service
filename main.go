package main

import (
	profilesgrpc "github.com/bombergame/auth-service/clients/profiles-service/grpc"
	"github.com/bombergame/common/logs"
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
}
