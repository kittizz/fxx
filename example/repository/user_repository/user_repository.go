package user_repository

import "github.com/kittizz/fxx/example/service/user_service"

type UserRepository struct {
	userService *user_service.UserService
}

func NewUserRepository(userService *user_service.UserService) *UserRepository {
	return &UserRepository{
		userService: userService,
	}
}

func (_this UserRepository) GetUser() string {
	return _this.userService.MockUser()
}
