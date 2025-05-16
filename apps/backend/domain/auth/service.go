package auth

import (
	"context"
	"tsm/domain/user"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(context.Context, LoginPayload) (*LoginResult, error)
	GetCurrentUser(context.Context, string) (*user.UserData, error)
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

func (service *authService) GetCurrentUser(ctx context.Context, authorization string) (*user.UserData, error) {
	if authorization == "" {
		return nil, nil
	}

	// TODO: validate and decode the auth token to get the user id from the payload
	userId, err := uuid.Parse(authorization)
	if err != nil {
		return nil, err
	}

	return service.user.GetById(ctx, userId)
}
