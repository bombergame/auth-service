package domains

//go:generate easyjson

//easyjson:json
type Session struct {
	ProfileID    int64  `json:"profile_id"`
	UserAgent    string `json:"-"`
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}
