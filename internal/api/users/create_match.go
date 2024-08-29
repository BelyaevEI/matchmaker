package users

import (
	"context"
	"log"
)

// CreateMatch Create match for users
func (i *Implementation) CreateMatch(ctx context.Context) error {

	// Find users
	users, err := i.userService.FindPalyers(ctx)
	if err != nil {
		log.Printf("not found users for match, %v", err)
	}

	if len(users) != 0 {
		// Print about new group
		i.userService.PrintNewGroup(users)
	}
	return nil
}
