package utils

import (
	"errors"
	"math"
	"time"

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

func DistanceMin(a, b model.User) float64 {
	return math.Sqrt(
		math.Pow(float64(a.Skill-b.Skill), 2) +
			math.Pow(float64(a.Latency-b.Latency), 2))
}

// Need use generic
func InfoSkill(users []model.User) (min, max int, avg float64) {
	if len(users) == 0 {
		return 0, 0, 0.0
	}

	max = math.MinInt
	min = math.MaxInt

	var sum int

	for _, user := range users {

		if int(user.Skill) > max {
			max = int(user.Skill)
		}

		if int(user.Skill) < min {
			min = int(user.Skill)
		}

		sum += int(user.Skill)
	}

	avg = float64(sum) / float64(len(users))

	return max, min, avg
}

func InfoLatency(users []model.User) (min, max int, avg float64) {
	if len(users) == 0 {
		return 0, 0, 0.0
	}

	max = math.MinInt
	min = math.MaxInt

	var sum int

	for _, user := range users {

		if int(user.Latency) > max {
			max = int(user.Skill)
		}

		if int(user.Latency) < min {
			min = int(user.Skill)
		}

		sum += int(user.Latency)
	}

	avg = float64(sum) / float64(len(users))

	return max, min, avg
}

func InfoTime(users []model.User) (maxDuration time.Duration, minDuration time.Duration, avgDuration time.Duration) {
	if len(users) == 0 {
		return 0, 0, 0
	}

	now := time.Now()

	maxDuration = -time.Duration(math.MaxInt64)
	minDuration = time.Duration(math.MaxInt64)

	var totalDuration time.Duration
	for _, user := range users {
		duration := now.Sub(user.TimeQueue)

		if duration > maxDuration {
			maxDuration = duration
		}

		if duration < minDuration {
			minDuration = duration
		}

		totalDuration += duration
	}

	avgDuration = totalDuration / time.Duration(len(users))

	return maxDuration, minDuration, avgDuration

}
