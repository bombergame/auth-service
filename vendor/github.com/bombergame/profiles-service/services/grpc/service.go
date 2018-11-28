package profilesgrpc

import (
	"context"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/grpc"
	"github.com/bombergame/profiles-service/config"
	"github.com/bombergame/profiles-service/repositories"
)

type Service struct {
	grpc.Service
	config     ServiceConfig
	components ServiceComponents
}

type ServiceConfig struct {
	grpc.ServiceConfig
}

type ServiceComponents struct {
	grpc.ServiceComponents
	ProfileRepository repositories.ProfileRepository
}

func NewService(cf ServiceConfig, cp ServiceComponents) *Service {
	cf.Host, cf.Port = consts.EmptyString, config.GrpcPort

	srv := &Service{
		config:     cf,
		components: cp,
		Service: *grpc.NewService(
			cf.ServiceConfig,
			cp.ServiceComponents,
		),
	}

	RegisterProfilesServiceServer(srv.Server, srv)

	return srv
}

func (srv *Service) IncProfileScore(ctx context.Context, req *ProfileID) (*Void, error) {
	return &Void{}, nil //TODO
}

func (srv *Service) GetProfileIDByCredentials(ctx context.Context, req *Credentials) (*ProfileID, error) {
	id, err := srv.components.ProfileRepository.FindIDByCredentials(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	profileID := &ProfileID{
		Value: *id,
	}

	return profileID, nil
}
