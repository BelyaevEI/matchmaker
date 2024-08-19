package model

import "time"

type User struct {
	Name      string    `db:"name"`
	Skill     int32     `db:"skill"`
	Latency   int32     `db:"latency"`
	TimeQueue time.Time `db:"created_at"`
}
