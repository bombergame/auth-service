package config

import (
	"github.com/bombergame/common/args"
	"github.com/bombergame/common/env"
)

var (
	HttpPort = args.GetString("http_port", "80")
	GrpcPort = args.GetString("grpc_port", "3000")

	ProfilesServiceGrpcAddress = env.GetVar(
		"PROFILES_SERVICE_GRPC_ADDRESS",
		"profiles-service:3000",
	)
)
