package container

import (
	"frogsmash/internal/app/user/factories"
	"frogsmash/internal/app/user/repos"
	"frogsmash/internal/app/user/services"
	"frogsmash/internal/config"
)

type User struct {
	UserService services.UserService
}

func NewUser(cfg *config.Config) *User {
	userFactory := factories.NewUserFactory()
	userRepo := repos.NewUserRepo()
	userService := services.NewUserService(userFactory, userRepo)

	return &User{
		UserService: userService,
	}
}
