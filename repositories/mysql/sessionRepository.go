package mysql

import (
	"github.com/bombergame/auth-service/domains"
	"github.com/bombergame/common/errs"
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
	statement, err := r.conn.db.Prepare(
		`REPLACE INTO session(profile_id,user_agent,refresh_token) VALUES(?,?,?);`,
	)
	if err != nil {
		return errs.NewServiceError(err)
	}

	_, err = statement.Exec(session.ProfileID, session.UserAgent, session.RefreshToken)
	if err != nil {
		return r.wrapError(err)
	}

	return nil
}

func (r *SessionRepository) RefreshSession(session domains.Session) error {
	statement, err := r.conn.db.Prepare(
		`SELECT EXISTS(SELECT * FROM session
         	WHERE profile_id = ? AND user_agent = ? AND refresh_token = ?
      	) AS token_exists;`,
	)
	if err != nil {
		return errs.NewServiceError(err)
	}

	row := statement.QueryRow(session.ProfileID, session.UserAgent, session.RefreshToken)

	var tokenExists int8
	if err := row.Scan(&tokenExists); err != nil {
		return errs.NewServiceError(err)
	}

	if tokenExists != 1 {
		return errs.NewNotAuthorizedError()
	}

	return nil
}

func (r *SessionRepository) DeleteSession(session domains.Session) error {
	statement, err := r.conn.db.Prepare(
		`DELETE FROM session WHERE profile_id = ? AND user_agent = ?;`,
	)
	if err != nil {
		return errs.NewServiceError(err)
	}

	_, err = statement.Exec(session.ProfileID, session.UserAgent)
	if err != nil {
		return r.wrapError(err)
	}

	return nil
}

func (r *SessionRepository) DeleteAllSessions(profileID int64) error {
	statement, err := r.conn.db.Prepare(
		`DELETE FROM session WHERE profile_id = ?;`,
	)
	if err != nil {
		return errs.NewServiceError(err)
	}

	_, err = statement.Exec(profileID)
	if err != nil {
		return r.wrapError(err)
	}

	return nil
}

func (r *SessionRepository) wrapError(err error) error {
	return errs.NewServiceError(err)
}
