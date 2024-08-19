package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/matchmaker/internal/utils"
)

func (s *service) PrintNewGroup(users []model.User) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.countGroup += 1

	// Text for log
	numberGroup := "Последовательный номер группы - " + strconv.Itoa(int(s.countGroup)) + "\n"

	skillMin, skillMax, skillAvg := utils.InfoSkill(users)
	skillGroup := fmt.Sprintf("Skill group - %v, %v, %v \n", skillMin, skillMax, skillAvg)

	latencyMin, latencyMax, latnecyAvg := utils.InfoLatency(users)
	latencyGroup := fmt.Sprintf("Latency group - %v, %v, %v \n", latencyMin, latencyMax, latnecyAvg)

	timeMin, timeMax, timeAvg := utils.InfoTime(users)
	timeGroup := fmt.Sprintf("Time spent in queue - %v, %v, %v \n", timeMin, timeMax, timeAvg)

	var names string
	for _, user := range users {
		names += user.Name + ","
	}
	namesGroup := "Список имен игроков: " + names + "\n"

	message := numberGroup + skillGroup + latencyGroup + timeGroup + namesGroup
	log.Println(message)
}
