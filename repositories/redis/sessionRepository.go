package redis

import (
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/common/errs"
	"strconv"
	"strings"
	"time"
)

const (
	keysSeparator = "|"
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
	kp := r.genKeysID(session.ProfileID)
	k, v := r.genKey(session), r.genValue(session)

	kpCmd := r.conn.client.Append(kp, keysSeparator+k)
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
		return r.wrapError(err)
	}
	if v != cmd.Val() {
		return errs.NewNotAuthorizedError()
	}
	return nil
}

func (r *SessionRepository) DeleteSession(session domains.Session) error {
	k := r.genKey(session)
	cmd := r.conn.client.Del(k)
	if err := cmd.Err(); err != nil {
		return errs.NewServiceError(err)
	}
	return nil
}

func (r *SessionRepository) DeleteAllSessions(id int64) error {
	kp := r.genKeysID(id)

	kpCmd := r.conn.client.Get(kp)
	if err := kpCmd.Err(); err != nil {
		return r.wrapError(err)
	}

	keys := strings.Split(kpCmd.Val(), keysSeparator)
	for _, k := range keys {
		cmd := r.conn.client.Del(k)
		if err := cmd.Err(); err != nil {
			return errs.NewServiceError(err)
		}
	}

	cmd := r.conn.client.Del(kp)
	if err := cmd.Err(); err != nil {
		return errs.NewServiceError(err)
	}

	return nil
}

func (r *SessionRepository) genKey(session domains.Session) string {
	return session.AuthToken
}

func (r *SessionRepository) genValue(session domains.Session) string {
	return strings.Join([]string{
		strconv.FormatInt(session.ProfileID, 10), session.UserAgent,
	}, ";")
}

func (r *SessionRepository) genKeysID(profileID int64) string {
	return strconv.FormatInt(profileID, 10)
}

func (r *SessionRepository) wrapError(err error) error {
	if err.Error() == "redis: nil" {
		return errs.NewNotAuthorizedError()
	}
	return errs.NewServiceError(err)
}
