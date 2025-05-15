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

type UserService struct {
	pool *domain.DatabasePool
}

func NewService(pool *domain.DatabasePool) UserService {
	return UserService{pool}
}

func (service *UserService) Create(ctx context.Context, payload UserCreateData) (*UserData, error) {
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

func (service *UserService) GetById(ctx context.Context, id uuid.UUID) (*UserData, error) {
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

func (service *UserService) GetByEmail(ctx context.Context, email string) (*UserData, error) {
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

func (service *UserService) GetByEmailAndPassword(ctx context.Context, email string, password string) (*UserData, error) {
	queries, release, err := service.pool.Acquire(ctx)
	defer release()

	if err != nil {
		return nil, err
	}

	user, err := queries.GetUserByEmail(ctx, email)

	if err == domain.ErrNoRows {
		return nil, ErrIncorrectUsernamePassword
	}

	if err != nil {
		return nil, err
	}

	validPassword := crypto.VerifyPassword(password, user.Password.String)

	if !validPassword {
		return nil, ErrIncorrectUsernamePassword
	}

	data := &UserData{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}

	return data, nil
}
