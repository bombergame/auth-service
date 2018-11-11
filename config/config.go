package config

import (
	"github.com/bombergame/common/args"
)

var (
	HttpPort = args.GetString("http_port", "80")
	GrpcPort = args.GetString("grpc_port", "3000")
)
