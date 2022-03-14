package main

import (
	"fmt"
	"os"

	"go.uber.org/fx"

	"github.com/kittizz/fxx"
	"github.com/kittizz/fxx/example/repository"
	"github.com/kittizz/fxx/example/repository/order_repository"
	"github.com/kittizz/fxx/example/repository/user_repository"
	"github.com/kittizz/fxx/example/service/user_service"
)

func main() {
	fx.New(
		fxx.MultiProvide(
			repository.NewRepository,
			user_service.NewUserService,
		),
		fx.Invoke(
			func(
				userRepo *user_repository.UserRepository,
				orderRepo *order_repository.OrderRepository,
			) {
				fmt.Println("Loadin :", userRepo.GetUser())
				for k, v := range orderRepo.List() {
					fmt.Println("My Order >", "["+k+"]", v)
				}
				os.Exit(0)
			},
		),
	).Run()
}
