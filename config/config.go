package config

import (
	"github.com/bombergame/common/args"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/env"
)

var (
	HttpPort = args.GetString("http_port", "80")
	GrpcPort = args.GetString("grpc_port", "3000")

	SessionsStorageAddress  = env.GetVar("SESSIONS_STORAGE_ADDRESS", "127.0.0.1:6379")
	SessionsStoragePassword = env.GetVar("SESSIONS_STORAGE_PASSWORD", consts.EmptyString)

	ProfilesServiceGrpcAddress = env.GetVar(
		"PROFILES_SERVICE_GRPC_ADDRESS",
		"profiles-service:3000",
	)
)
