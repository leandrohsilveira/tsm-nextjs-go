package user

import "tsm/database"

type UserData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateData struct {
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Password string            `json:"password"`
	Role     database.UserRole `json:"role"`
}
