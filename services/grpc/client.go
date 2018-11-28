package authgrpc

import (
	"github.com/bombergame/common/grpc"
)

type Client struct {
	grpc.Service
}

type ClientConfig struct {
	grpc.ClientConfig
}

type ClientComponents struct {
	grpc.ClientComponents
}
