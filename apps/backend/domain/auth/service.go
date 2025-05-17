package auth

import (
	"context"
	"tsm/domain"
	"tsm/domain/user"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type AuthService interface {
	Login(context.Context, LoginPayload) (*LoginResult, error)
	GetCurrentUser(context.Context, LoginInfoPayload) (*user.UserData, error)
}

type authService struct {
	user user.UserService
}

func NewService(userService user.UserService) AuthService {
	return &authService{user: userService}
}

func (service *authService) Login(ctx context.Context, payload LoginPayload) (*LoginResult, error) {
	data, err := service.user.GetByEmailAndPassword(ctx, payload.Username, payload.Password)

	if err != nil {
		return nil, err
	}

	return &LoginResult{Token: data.ID, RefreshToken: ""}, nil
}

func (service *authService) GetCurrentUser(ctx context.Context, payload LoginInfoPayload) (*user.UserData, error) {
	authorization := payload.Token

	// TODO: validate and decode the auth token to get the user id from the payload
	userId, err := uuid.Parse(authorization)
	if err != nil {
		return nil, err
	}

	data, err := service.user.GetById(ctx, userId)
	if err == domain.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	log.Ctx(ctx).Info().Str("email", data.Email).Str("ID", data.ID).Msg("Get current user info")

	return data, nil
}
