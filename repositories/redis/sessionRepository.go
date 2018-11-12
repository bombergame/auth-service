package redis

import (
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/common/errs"
	"strconv"
	"strings"
	"time"
)

type SessionRepository struct {
	conn       *Connection
	expireTime time.Duration
}

func NewSessionRepository(conn *Connection) *SessionRepository {
	t, err := strconv.Atoi(config.SessionExpireTime)
	if err != nil {
		panic(err)
	}

	return &SessionRepository{
		conn:       conn,
		expireTime: time.Duration(t) * time.Second,
	}
}

func (r *SessionRepository) AddSession(session domains.Session) error {
	kp := r.genKeyPart(session.ProfileID)
	k, v := r.genKey(session), r.genValue(session)

	kpCmd := r.conn.client.Append(kp, k)
	if err := kpCmd.Err(); err != nil {
		return errs.NewServiceError(err)
	}

	kCmd := r.conn.client.Set(k, v, r.expireTime)
	if err := kCmd.Err(); err != nil {
		return errs.NewServiceError(err)
	}

	return nil
}

func (r *SessionRepository) CheckSession(session domains.Session) error {
	k, v := r.genKey(session), r.genValue(session)
	cmd := r.conn.client.Get(k)
	if err := cmd.Err(); err != nil {
		return wrapError(err)
	}
	if v != cmd.String() {
		return errs.NewNotAuthorizedError()
	}
	return nil
}

func (r *SessionRepository) DeleteAllSessionsByProfileID(id int64) error {
	return nil //TODO
}

func (r *SessionRepository) genKey(session domains.Session) string {
	return strings.Join([]string{
		strconv.FormatInt(session.ProfileID, 10),
		session.UserAgent}, ";",
	)
}

func (r *SessionRepository) genKeyPart(id int64) string {
	return strconv.FormatInt(id, 10)
}

func (r *SessionRepository) genValue(session domains.Session) string {
	return session.AuthToken
}

func wrapError(err error) error {
	return errs.NewServiceError(err)
}
