package redis

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

func (r *SessionRepository) CreateSession(session domains.Session) error {
	return nil //TODO
}

func (r *SessionRepository) GetProfileIDByAuthToken(token string) (*int64, error) {
	return nil, nil //TODO
}

func (r *SessionRepository) DeleteAllSessionsByProfileID(id int64) error {
	return nil //TODO
}
