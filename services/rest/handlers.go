package rest

import (
	"github.com/bombergame/auth-service/clients/profiles-service/grpc"
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/auth-service/utils"
	"github.com/bombergame/common/auth"
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

	//id, err := srv.config.ProfilesGrpc.GetProfileIDByCredentials(
	//	&profilesgrpc.Credentials{
	//		Username: cr.Username,
	//		Password: cr.Password,
	//	},
	//)
	//if err != nil {
	//	srv.writeErrorWithBody(w, err)
	//	return
	//}

	//token, err := srv.config.TokenManager.CreateToken(
	//	utils.UserInfo{
	//		ProfileID: id.Value,
	//		UserAgent: userAgent,
	//	},
	//)
	//if err != nil {
	//	srv.writeErrorWithBody(w, err)
	//	return
	//}
	//
	//err = srv.config.SessionRepository.AddSession(
	//	domains.Session{
	//		ProfileID: id.Value,
	//		UserAgent: userAgent,
	//		AuthToken: token,
	//	},
	//)
	//if err != nil {
	//	srv.writeErrorWithBody(w, err)
	//	return
	//}
	//
	//srv.writeOkWithBody(w, Session{
	//	ProfileID: id.Value,
	//	AuthToken: token,
	//})
}

func (srv *Service) refreshSession(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) deleteSession(w http.ResponseWriter, r *http.Request) {
	authToken, err := srv.ReadAuthToken(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	userInfo, err := srv.config.TokenManager.ParseToken(authToken)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	err = srv.config.SessionRepository.DeleteSession(
		domains.Session{
			AuthToken: authToken,
			ProfileID: userInfo.ProfileID,
			UserAgent: userInfo.UserAgent,
		},
	)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOk(w)
}
