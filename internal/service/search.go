package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/matchmaker/internal/utils"
)

// Method add user for search game
func (s *service) SearchMatch(ctx context.Context, body []byte) error {

	var user model.User

	if err := json.Unmarshal(body, &user); err != nil {
		log.Printf("deserealizating body is failed")
		return err
	}

	if err := utils.CheckNotEmpty(user); err != nil {
		log.Printf("user parameters is empty")
		return err
	}

	if err := s.userRepository.AddUserToPool(ctx, user); err != nil {
		log.Printf("add user to pool is failed")
		return err
	}

	return nil
}
