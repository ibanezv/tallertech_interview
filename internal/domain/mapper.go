package domain

import "github.com/ibanezv/tallertech_interview/internal/infrastructure/repo"

func ToDbEvent(event *Event) *repo.EventDb {
	return &repo.EventDb{
		Id:          event.Id,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		CreateAt:    event.CreateAt,
	}
}

func ToEvent(eventDb *repo.EventDb) *Event {
	return &Event{
		Id:          eventDb.Id,
		Description: eventDb.Description,
		StartTime:   eventDb.StartTime,
		EndTime:     eventDb.EndTime,
		CreateAt:    eventDb.CreateAt,
	}
}

func ToEventList(eventDbList []repo.EventDb) []Event {
	list := make([]Event, len(eventDbList))
	for i := 0; i < len(eventDbList); i++ {
		list = append(list, Event{
			Id:          eventDbList[i].Id,
			Description: eventDbList[i].Description,
			StartTime:   eventDbList[i].StartTime,
			EndTime:     eventDbList[i].EndTime,
			CreateAt:    eventDbList[i].CreateAt,
		})
	}
	return list
}
