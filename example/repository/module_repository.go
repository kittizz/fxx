package repository

import (
	"github.com/kittizz/fxx"
	"github.com/kittizz/fxx/example/repository/order_repository"
	"github.com/kittizz/fxx/example/repository/user_repository"
)

var NewRepository = fxx.Modules(
	user_repository.NewUserRepository,
	order_repository.NewOrderRepository,
)
