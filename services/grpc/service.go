package grpc

import (
	"context"
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/repositories"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/grpc"
)

type Service struct {
	grpc.Service
	config     Config
	components Components
}

type Config struct {
	grpc.Config
}

type Components struct {
	grpc.Components
	sessionRepository repositories.SessionRepository
}

func NewService(cf Config, cp Components) *Service {
	cf.Host, cf.Port = consts.EmptyString, config.GrpcPort

	srv := &Service{
		config:     cf,
		components: cp,
		Service: *grpc.NewService(
			cf.Config,
			cp.Components,
		),
	}

	RegisterAuthServiceServer(srv.Server, srv)

	return srv
}

func (srv *Service) DeleteAllSessions(context.Context, *ProfileID) (*Void, error) {
	return nil, nil //TODO
}
