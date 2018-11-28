package authrest

import (
	"github.com/bombergame/common/auth"
	"strconv"
	"time"
)

func (srv *Service) formatInt64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func (srv *Service) parseInt64(v string) int64 {
	iV, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return iV
}

func (srv *Service) getTokenExpireTime() time.Time {
	return time.Now().Add(auth.DefaultTokenValidDuration)
}
