package authrest

import (
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/repositories"
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/rest"
	"github.com/bombergame/profiles-service/services/grpc"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Service struct {
	rest.Service
	config     Config
	components Components
}

type Config struct {
	rest.Config
}

type Components struct {
	rest.Components
	AuthTokenManager    auth.TokenManager
	RefreshTokenManager auth.TokenManager
	SessionRepository   repositories.SessionRepository
	ProfilesClient      *profilesgrpc.Client
}

func NewService(cf Config, cpn Components) *Service {
	cf.Host, cf.Port = consts.EmptyString, config.HttpPort

	srv := &Service{
		Service: *rest.NewService(
			cf.Config,
			cpn.Components,
		),
		config:     cf,
		components: cpn,
	}

	mx := http.NewServeMux()
	mx.Handle("/auth/session", handlers.MethodHandler{
		http.MethodPost:   http.HandlerFunc(srv.createSession),
		http.MethodPatch:  http.HandlerFunc(srv.refreshSession),
		http.MethodDelete: http.HandlerFunc(srv.deleteSession),
	})

	mx.Handle("/metrics", promhttp.Handler())

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
