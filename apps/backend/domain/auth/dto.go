package auth

import "tsm/domain/user"

type LoginPayload struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type LoginResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type LoginInfo struct {
	Data user.UserData `json:"data"`
}
