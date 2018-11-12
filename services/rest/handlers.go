package rest

import (
	profilesgrpc "github.com/bombergame/auth-service/clients/profiles-service/grpc"
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/auth-service/utils"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"net/http"
)

func (srv *Service) createSession(w http.ResponseWriter, r *http.Request) {
	userAgent, err := srv.readUserAgent(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	var cr Credentials
	if err := srv.readRequestBody(&cr, r); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if err := cr.Validate(); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	id, err := srv.config.ProfilesGrpc.GetProfileIDByCredentials(
		&profilesgrpc.Credentials{
			Username: cr.Username,
			Password: cr.Password,
		},
	)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	token, err := srv.config.TokenManager.CreateToken(
		utils.UserInfo{
			ProfileID: id.Value,
			UserAgent: userAgent,
		},
	)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	err = srv.config.SessionRepository.AddSession(
		domains.Session{
			ProfileID: id.Value,
			UserAgent: userAgent,
			AuthToken: token,
		},
	)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOkWithBody(w, Session{
		ProfileID: id.Value,
		AuthToken: token,
	})
}

func (srv *Service) deleteSession(w http.ResponseWriter, r *http.Request) {
	_, err := srv.readUserAgent(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	_, err = srv.readAuthToken(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOk(w)
}

func (srv *Service) readUserAgent(r *http.Request) (string, error) {
	return srv.readHeaderString("X-User-Agent", r)
}

func (srv *Service) readAuthToken(r *http.Request) (string, error) {
	bearer, err := srv.readHeaderString("Authorization", r)
	if err != nil {
		return bearer, err
	}

	n := len("Bearer ")
	if len(bearer) <= n {
		return consts.EmptyString, errs.NewInvalidFormatError("wrong authorization token")
	}

	return bearer[n:], nil
}
