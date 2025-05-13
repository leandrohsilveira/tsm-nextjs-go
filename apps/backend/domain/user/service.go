package user

import (
	"context"
	"tsm/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) UserService {
	return UserService{pool}
}

func (service *UserService) GetById(ctx context.Context, id uuid.UUID) (UserData, error) {
	conn, err := service.pool.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return UserData{}, err
	}

	queries := database.New(conn)

	user, err := queries.GetUserById(ctx, id)

	if err != nil {
		return UserData{}, err
	}

	data := UserData{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}

	return data, nil
}

func (service *UserService) GetByEmailAndPassword(ctx context.Context, email string, password string) (UserData, error) {
	conn, err := service.pool.Acquire(ctx)

	defer conn.Release()

	if err != nil {
		return UserData{}, err
	}

	queries := database.New(conn)

	user, err := queries.GetUserByEmail(ctx, email)

	if err != nil {
		return UserData{}, err
	}

	// TODO: decode and check user password

	data := UserData{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}

	return data, nil

}
