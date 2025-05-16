package main

import (
	"context"
	"os"
	"tsm/database"
	"tsm/domain"
	"tsm/domain/user"

	"github.com/gofiber/fiber/v2/log"
)

func Seed(ctx context.Context, pool domain.DatabasePool) error {
	activate := ShouldSeed()

	if !activate {
		return nil
	}

	log.Info("Seeding database...")

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
		log.Infof("Admin user %s already exists (ID %s), seed will be aborted", username, data.ID)
		return nil
	}

	data, err = service.Create(ctx, user.UserCreateData{
		Name:     username,
		Email:    username,
		Password: password,
		Role:     database.UserRoleAdminUser,
	})

	if err == nil {
		log.Infof("Admin user %s created with ID: %s", username, data.ID)
	}

	return err
}

func ShouldSeed() bool {
	args := os.Args[1:]

	activate := false
	for _, arg := range args {
		if arg == "--seed" {
			activate = true
		}
	}

	return activate
}
