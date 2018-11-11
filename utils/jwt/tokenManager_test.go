package jwt

import (
	"github.com/bombergame/auth-service/utils"
	"testing"
)

func TestTokenManagerUnit(t *testing.T) {
	manager := NewTokenManager()
	if manager == nil {
		t.Error("token manager not created")
	}

	info := utils.UserInfo{
		ProfileID:  100,
		UserAgent: "some-user-agent",
	}

	token, err := manager.CreateToken(info)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	pInfo, err := manager.ParseToken(token)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if info != *pInfo {
		t.Error("tokens differ")
	}

	pInfo, err = manager.ParseToken("some_invalid_token")
	if err == nil {
		t.Error("invalid token parsed")
	}
}
