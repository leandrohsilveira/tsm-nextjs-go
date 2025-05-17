package auth

import "tsm/domain/user"

type LoginPayload struct {
	Username string `form:"username" json:"username" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
}

type LoginResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type LoginInfo struct {
	Data user.UserData `json:"data"`
}

type LoginInfoPayload struct {
	Token string `reqHeader:"authorization" validate:"required"`
}
