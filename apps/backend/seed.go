package main

import (
	"context"
	"os"
	"tsm/database"
	"tsm/domain"
	"tsm/domain/user"

	"github.com/rs/zerolog/log"
)

func Seed(ctx context.Context, pool domain.DatabasePool) error {
	if !domain.ShouldSeed {
		return nil
	}

	log.Info().Msg("Seeding database")

	username, isUsernameSet := os.LookupEnv("ADMIN_USERNAME")
	if !isUsernameSet {
		username = "admin@email.com"
	}

	password, isPasswordSet := os.LookupEnv("ADMIN_PASSWORD")
	if !isPasswordSet {
		password = "123456"
	}

	service := user.NewService(pool)

	data, err := service.GetByEmail(ctx, username)

	if err != nil {
		return err
	}

	if data != nil {
		log.Info().
			Str("email", username).
			Str("ID", data.ID).
			Msgf("Admin user already exists, admin user seed will be aborted")

		return nil
	}

	data, err = service.Create(ctx, user.UserCreateData{
		Name:     username,
		Email:    username,
		Password: password,
		Role:     database.UserRoleAdminUser,
	})

	if err == nil {
		log.Info().
			Str("email", username).
			Str("ID", data.ID).
			Msgf("Admin user created")
	}

	return err
}
