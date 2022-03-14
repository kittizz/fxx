package order_repository

import "github.com/kittizz/fxx/example/service/user_service"

type OrderRepository struct {
	userService *user_service.UserService
}

func NewOrderRepository(userService *user_service.UserService) *OrderRepository {
	return &OrderRepository{
		userService: userService,
	}
}
func (_this OrderRepository) List() map[string]string {
	return map[string]string{
		_this.userService.MockUser(): "IPHONE X",
	}
}
