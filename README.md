# 🦜 Fxx [![GoDoc][doc-img]][doc]

An add on for [Fx](https://github.com/uber-go/fx):

* add `func MultiProvide()  fx.Option`.

## Installation

We recommend locking to [SemVer](http://semver.org/) range `^1` using [go mod](https://github.com/golang/go/wiki/Modules):

```shell
go get go.uber.org/fx@v1
go get github.com/kittizz/fxx
```

## Example
```go
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
....

package repository

var NewRepository = fxx.Modules(
	user_repository.NewUserRepository,
	order_repository.NewOrderRepository,
)
....

package user_service

func NewUserService() *UserService {
	return &UserService{}
}



```
  

[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://godoc.org/github.com/kittizz/fxx

 