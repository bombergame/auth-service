package jwt

import (
	"github.com/bombergame/auth-service/utils"
	"github.com/bombergame/common/env"
	"github.com/bombergame/common/errs"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
)

const (
	DefaultKeyLength = 64
)

type TokenManager struct {
	key []byte
}

func NewTokenManager() *TokenManager {
	key := env.GetVar("TOKEN_SIGN_KEY", "")
	if key == "" {
		key = generateKey()
	}

	return &TokenManager{
		key: []byte(key),
	}
}

func (m *TokenManager) CreateToken(info utils.UserInfo) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"profile_id": info.ProfileID,
		"user_agent": info.UserAgent,
	})
	return t.SignedString(m.key)
}

func (m *TokenManager) ParseToken(token string) (*utils.UserInfo, error) {
	t, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewInvalidFormatError("wrong signing method")
		}
		return m.key, nil
	})

	if err != nil {
		return nil, errs.NewInvalidFormatError(err.Error())
	}

	if !t.Valid {
		return nil, errs.NewInvalidFormatError("wrong token")
	}

	errWrongClaims := errs.NewInvalidFormatError("wrong claims")

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errWrongClaims
	}

	vID, ok := claims["profile_id"]
	if !ok {
		return nil, errWrongClaims
	}

	profileID, ok := vID.(float64)
	if !ok {
		return nil, errWrongClaims
	}

	vUserAgent, ok := claims["user_agent"]
	if !ok {
		return nil, errWrongClaims
	}

	userAgent := vUserAgent.(string)
	if !ok {
		return nil, errWrongClaims
	}

	info := &utils.UserInfo{
		ProfileID: int64(profileID),
		UserAgent: userAgent,
	}

	return info, nil
}

func generateKey() string {
	key := make([]rune, DefaultKeyLength)
	runes := []rune(`abcdefghijklmnopqrstuvwxyz1234567890@#$^&*()_-=+`)

	n := len(runes)
	for i := range key {
		key[i] = runes[rand.Intn(n)]
	}

	return string(key)
}
