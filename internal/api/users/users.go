package users

import "github.com/BelyaevEI/matchmaker/internal/service"

type Implementation struct {
	userService service.UserServicer
}

func NewImplementation(userService service.UserServicer) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
