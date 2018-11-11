package rest

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
)

// easyjson:json
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cr Credentials) Validate() error {
	if cr.Username == consts.EmptyString {
		return errs.NewInvalidFormatError("empty username")
	}
	if cr.Password == consts.EmptyString {
		return errs.NewInvalidFormatError("empty password")
	}
	return nil
}

// easyjson:json
type Session struct {
	ProfileID int64  `json:"profile_id"`
	AuthToken string `json:"auth_token"`
}
