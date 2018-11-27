package rest

import (
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/repositories"
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/rest"
	"github.com/gorilla/handlers"
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
	authTokenManager    auth.TokenManager
	refreshTokenManager auth.TokenManager
	sessionRepository   repositories.SessionRepository
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
		http.MethodPost:   srv.WithAuth(http.HandlerFunc(srv.createSession)),
		http.MethodPatch:  srv.WithAuth(http.HandlerFunc(srv.refreshSession)),
		http.MethodDelete: srv.WithAuth(http.HandlerFunc(srv.deleteSession)),
	})

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
