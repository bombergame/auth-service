package authgrpc

import (
	"context"
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/auth-service/repositories"
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/logs"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	config *Config
	server *grpc.Server
}

type Config struct {
	Logger            *logs.Logger
	TokenManager      auth.TokenManager
	SessionRepository repositories.SessionRepository
}

func NewService(c *Config) *Service {
	srv := &Service{
		config: c,
	}

	grpcSrv := grpc.NewServer()
	RegisterAuthServiceServer(grpcSrv, srv)

	srv.server = grpcSrv

	return srv
}

func (srv *Service) Run() error {
	addr := ":" + config.GrpcPort

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv.config.Logger.Info("grpc service running on: " + addr)
	return srv.server.Serve(lis)
}

func (srv *Service) Shutdown() error {
	srv.config.Logger.Info("grpc service shutdown initialized")
	srv.server.GracefulStop()
	return nil
}

func (srv *Service) GetProfileID(ctx context.Context, req *AuthInfo) (*ProfileID, error) {
	info, err := srv.config.TokenManager.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	err = srv.config.SessionRepository.CheckSession(
		domains.Session{
			ProfileID: info.ProfileID,
			UserAgent: info.UserAgent,
			AuthToken: req.Token,
		},
	)
	if err != nil {
		return nil, err
	}

	pID := &ProfileID{
		Value: info.ProfileID,
	}

	return pID, nil
}

func (srv *Service) DeleteAllSessions(ctx context.Context, req *ProfileID) (*Void, error) {
	id := req.Value

	err := srv.config.SessionRepository.DeleteAllSessions(id)
	if err != nil {
		return nil, err
	}

	return &Void{}, nil
}
