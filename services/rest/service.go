package rest

import (
	"context"
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/common/logs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Service struct {
	config *Config
	server http.Server
}

type Config struct {
	Logger *logs.Logger
}

func NewService(c *Config) *Service {
	srv := &Service{
		config: c,
		server: http.Server{
			Addr: ":" + config.HttpPort,
		},
	}

	mx := mux.NewRouter()

	mx.Handle("/session", handlers.MethodHandler{
		http.MethodPost:   http.HandlerFunc(srv.createSession),
		http.MethodDelete: http.HandlerFunc(srv.deleteSession),
	})

	srv.server.Handler = srv.withLogs(srv.withRecover(mx))

	return srv
}

func (srv *Service) Run() error {
	srv.config.Logger.Info("http service running on: " + srv.server.Addr)
	return srv.server.ListenAndServe()
}

func (srv *Service) Shutdown() error {
	srv.config.Logger.Info("http service shutdown initialized")
	return srv.server.Shutdown(context.TODO())
}
