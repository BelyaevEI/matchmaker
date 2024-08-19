package users

import (
	"context"
	"log"
)

// Create match for users
func (i *Implementation) CreateMatch(ctx context.Context) error {

	// Find users
	users, err := i.userService.FindPalyers(ctx)
	if err != nil {
		log.Printf("not find users for match, %v", err)
	}

	// Print about new group
	i.userService.PrintNewGroup(users)

	return nil
}
