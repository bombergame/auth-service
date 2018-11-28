package rest

import (
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/common/auth"
	"math/rand"
	"net/http"
)

func (srv *Service) createSession(w http.ResponseWriter, r *http.Request) {
	userAgent, err := srv.ReadUserAgent(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	var c auth.Credentials
	if err := srv.ReadRequestBody(&c, r); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if err := c.Validate(); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	id := rand.Int63() //TODO

	userInfo := auth.UserInfo{
		ProfileID: id,
		UserAgent: userAgent,
	}

	authToken, err := srv.components.AuthTokenManager.CreateToken(userInfo)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	refreshToken, err := srv.components.RefreshTokenManager.CreateToken(userInfo)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	session := domains.Session{
		ProfileID:    id,
		UserAgent:    userAgent,
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}

	err = srv.components.SessionRepository.AddSession(session)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOkWithBody(w, session)
}

func (srv *Service) refreshSession(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) deleteSession(w http.ResponseWriter, r *http.Request) {
	//TODO
}
