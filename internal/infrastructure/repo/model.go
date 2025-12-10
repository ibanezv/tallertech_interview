package repo

import "time"

type EventDb struct {
	Id          string
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	CreateAt    time.Time
}
