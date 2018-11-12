package repositories

import (
	"github.com/bombergame/auth-service/domains"
)

type SessionRepository interface {
	AddSession(session domains.Session) error
	CheckSession(session domains.Session) error
	DeleteAllSessionsByProfileID(id int64) error
}
