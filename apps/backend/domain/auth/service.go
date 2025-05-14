package auth

import (
	"context"
	"net/http"
	"tsm/domain/user"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	user user.UserService
}

func NewService(userService user.UserService) AuthService {
	return AuthService{user: userService}
}

func (service *AuthService) Login(ctx context.Context, payload LoginPayload) (*LoginResult, error) {
	data, err := service.user.GetByEmailAndPassword(ctx, payload.Username, payload.Password)

	if err != nil {
		return nil, err
	}

	return &LoginResult{Token: data.ID, RefreshToken: ""}, nil
}

func (service *AuthService) GetCurrentUser(req http.Request) (*user.UserData, error) {
	authorization := req.Header.Get("authorization")

	if authorization == "" {
		return nil, echo.ErrUnauthorized
	}

	// TODO: validate and decode the auth token to get the user id from the payload
	userId, err := uuid.Parse(authorization)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return service.user.GetById(req.Context(), userId)
}
