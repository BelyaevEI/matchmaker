package users

import "github.com/BelyaevEI/matchmaker/internal/service"

type Implementation struct {
	usersService service.UserServicer
}

func NewImplementation(usersService service.UserServicer) *Implementation {
	return &Implementation{
		usersService: usersService,
	}
}
