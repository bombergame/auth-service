package utils

type UserInfo struct {
	ProfileID int64
	UserAgent string
}

type TokenManager interface {
	CreateToken(info UserInfo) (string, error)
	ParseToken(token string) (*UserInfo, error)
}
