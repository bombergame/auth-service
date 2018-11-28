package authrest

import (
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/profiles-service/services/grpc"
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

	credentials := profilesgrpc.Credentials{
		Username: c.Username,
		Password: c.Password,
	}

	pID, err := srv.components.ProfilesClient.GetProfileIDByCredentials(&credentials)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	expireTime := srv.getTokenExpireTime()

	tokenInfo := auth.TokenInfo{
		ProfileID:  srv.formatInt64(pID.Value),
		UserAgent:  userAgent,
		ExpireTime: expireTime.Format(auth.ExpireTimeFormat),
	}

	authToken, err := srv.components.AuthTokenManager.CreateToken(tokenInfo)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	refreshToken, err := srv.components.RefreshTokenManager.CreateToken(tokenInfo)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	session := domains.Session{
		ProfileID:    pID.Value,
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
	userAgent, err := srv.ReadUserAgent(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	authToken, err := srv.ReadAuthToken(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	tokenInfo, err := srv.components.AuthTokenManager.ParseToken(authToken)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if tokenInfo.UserAgent != userAgent {
		err := errs.NewAccessDeniedError()
		srv.WriteErrorWithBody(w, err)
		return
	}

	var session domains.Session
	if err := srv.ReadRequestBody(&session, r); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	session.ProfileID = srv.parseInt64(tokenInfo.ProfileID)
	session.UserAgent = userAgent

	if err := srv.components.SessionRepository.RefreshSession(session); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	tokenInfo.ExpireTime = srv.getTokenExpireTime().Format(auth.ExpireTimeFormat)

	newAuthToken, err := srv.components.AuthTokenManager.CreateToken(*tokenInfo)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	session.AuthToken = newAuthToken
	srv.WriteOkWithBody(w, session)
}

func (srv *Service) deleteSession(w http.ResponseWriter, r *http.Request) {
	userAgent, err := srv.ReadUserAgent(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	authToken, err := srv.ReadAuthToken(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	tokenInfo, err := srv.components.AuthTokenManager.ParseToken(authToken)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if tokenInfo.UserAgent != userAgent {
		err := errs.NewAccessDeniedError()
		srv.WriteErrorWithBody(w, err)
		return
	}

	session := domains.Session{
		ProfileID: srv.parseInt64(tokenInfo.ProfileID),
		UserAgent: userAgent,
	}
	if err := srv.components.SessionRepository.DeleteSession(session); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOk(w)
}
