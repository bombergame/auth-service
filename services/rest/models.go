package rest

// easyjson:json
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// easyjson:json
type Session struct {
	ProfileID int64  `json:"profile_id"`
	AuthToken string `json:"auth_token"`
}
