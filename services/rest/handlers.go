package rest

import (
	profilesgrpc "github.com/bombergame/auth-service/clients/profiles-service/grpc"
	"github.com/bombergame/auth-service/utils"
	"net/http"
)

func (srv *Service) createSession(w http.ResponseWriter, r *http.Request) {
	userAgent, err := srv.readHeaderString("X-User-Agent", r)
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

	session := Session{
		ProfileID: id.Value,
		AuthToken: token,
	}

	srv.writeOkWithBody(w, session)
}

func (srv *Service) deleteSession(w http.ResponseWriter, r *http.Request) {
	//TODO
}
