package grpc

import (
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/logs"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	Config     Config
	Components Components
	Server     *grpc.Server
}

type Config struct {
	Host string
	Port string
}

type Components struct {
	Logger      *logs.Logger
	AuthManager auth.AuthenticationManager
}

func NewService(cf Config, cp Components) *Service {
	return &Service{
		Config:     cf,
		Components: cp,
		Server:     grpc.NewServer(),
	}
}

func (srv *Service) Run() error {
	addr := srv.Config.Host + ":" + srv.Config.Port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv.Logger().Info("grpc service running on: " + addr)
	return srv.Server.Serve(lis)
}

func (srv *Service) Shutdown() error {
	srv.Logger().Info("grpc service shutdown initialized")
	srv.Server.GracefulStop()
	return nil
}

func (srv *Service) Logger() *logs.Logger {
	return srv.Components.Logger
}