package mysql

import (
	"github.com/bombergame/auth-service/domains"
)

type SessionRepository struct {
	conn *Connection
}

func NewSessionRepository(conn *Connection) *SessionRepository {
	return &SessionRepository{
		conn: conn,
	}
}

func (r *SessionRepository) AddSession(session domains.Session) error {
	return nil //TODO
}

func (r *SessionRepository) CheckSession(session domains.Session) error {
	return nil //TODO
}

func (r *SessionRepository) DeleteSession(session domains.Session) error {
	return nil //TODO
}

func (r *SessionRepository) DeleteAllSessions(profileID int64) error {
	return nil //TODO
}
