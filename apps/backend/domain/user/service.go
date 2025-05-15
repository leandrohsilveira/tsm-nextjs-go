package user

import (
	"context"
	"errors"
	"tsm/crypto"
	"tsm/database"
	"tsm/domain"

	"github.com/google/uuid"
)

var ErrIncorrectUsernamePassword = errors.New("Incorrect username or password")

type UserService interface {
	Create(context.Context, UserCreateData) (*UserData, error)
	GetById(context.Context, uuid.UUID) (*UserData, error)
	GetByEmail(context.Context, string) (*UserData, error)
	GetByEmailAndPassword(context.Context, string, string) (*UserData, error)
}

type userService struct {
	pool domain.DatabasePool
}

func NewService(pool domain.DatabasePool) UserService {
	return &userService{pool}
}

func (service *userService) Create(ctx context.Context, payload UserCreateData) (*UserData, error) {
	password, err := crypto.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	queries, release, err := service.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer release()

	result, err := queries.CreateUser(ctx, database.CreateUserParams{
		Name:     payload.Name,
		Email:    payload.Email,
		Role:     payload.Role,
		Password: service.pool.Text(password),
	})
	if err != nil {
		return nil, err
	}

	data := &UserData{
		ID:    result.ID.String(),
		Name:  result.Name,
		Email: result.Email,
	}
	return data, nil
}

func (service *userService) GetById(ctx context.Context, id uuid.UUID) (*UserData, error) {
	queries, release, err := service.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer release()

	user, err := queries.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	data := &UserData{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}
	return data, nil
}

func (service *userService) GetByEmail(ctx context.Context, email string) (*UserData, error) {
	queries, release, err := service.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer release()

	user, err := queries.GetUserByEmail(ctx, email)
	if err == domain.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	data := &UserData{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}
	return data, nil
}

func (service *userService) GetByEmailAndPassword(ctx context.Context, email string, password string) (*UserData, error) {
	queries, release, err := service.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer release()

	user, err := queries.GetUserByEmail(ctx, email)
	if err == domain.ErrNoRows {
		return nil, ErrIncorrectUsernamePassword
	}
	if err != nil {
		return nil, err
	}

	matched, err := crypto.VerifyPassword(password, user.Password.String)
	if err != nil {
	}
	if !matched {
		return nil, ErrIncorrectUsernamePassword
	}

	data := &UserData{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}

	return data, nil
}
