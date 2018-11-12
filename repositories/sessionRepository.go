package repositories

import (
	"github.com/bombergame/auth-service/domains"
)

type SessionRepository interface {
	CreateSession(session domains.Session) error
	GetProfileIDByAuthToken(token string) (*int64, error)
	DeleteAllSessionsByProfileID(id int64) error
}
