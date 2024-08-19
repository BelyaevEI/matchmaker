package users

import "github.com/BelyaevEI/matchmaker/internal/service"

// Implementation api layer
type Implementation struct {
	userService service.UserServicer
}

// NewImplementation constructor
func NewImplementation(userService service.UserServicer) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
