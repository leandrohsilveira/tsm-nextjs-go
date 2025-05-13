package auth

import (
	"context"
	"errors"
	"net/http"
	"tsm/domain/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

var IncorrectUsernamePassword = errors.New("Incorrect username or password")

type AuthService struct {
	user user.UserService
}

func NewService(userService user.UserService) AuthService {
	return AuthService{user: userService}
}

func (service *AuthService) Login(ctx context.Context, payload LoginPayload) (LoginResult, error) {
	data, err := service.user.GetByEmailAndPassword(ctx, payload.Username, payload.Password)

	if err == pgx.ErrNoRows {
		return LoginResult{}, IncorrectUsernamePassword
	}

	// TODO: check password

	if err != nil {
		return LoginResult{}, err
	}

	return LoginResult{Token: data.ID, RefreshToken: ""}, nil
}

func (service *AuthService) GetCurrentUser(req http.Request) (user.UserData, error) {
	authorization := req.Header.Get("authorization")

	if authorization == "" {
		return user.UserData{}, echo.ErrUnauthorized
	}

	// TODO: validate and decode the auth token to get the user id from the payload
	userId, err := uuid.Parse(authorization)
	if err != nil {
		return user.UserData{}, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return service.user.GetById(req.Context(), userId)
}
