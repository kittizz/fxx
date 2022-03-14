package user_service

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (_this UserService) MockUser() string {
	return "TEST_UXXX"
}
