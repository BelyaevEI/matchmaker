package utils

import (
	"errors"

	"github.com/BelyaevEI/matchmaker/internal/model"
)

func CheckNotEmpty(user model.User) error {

	if len(user.Name) == 0 {
		return errors.New("user name is empty")
	}

	if user.Skill == 0 {
		return errors.New("user skill is empty")
	}

	if user.Latency == 0 {
		return errors.New("user latency is empty")
	}

	return nil
}
