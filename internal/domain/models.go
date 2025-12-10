package domain

import "time"

type Event struct {
	Id          string    `json:"id"`
	Title       string    `json:"title,required`
	Description string    `json:"description,required"`
	StartTime   time.Time `json:"start_time,required"`
	EndTime     time.Time `json:"end_time,required"`
	CreateAt    time.Time `json:"create_at"`
}
