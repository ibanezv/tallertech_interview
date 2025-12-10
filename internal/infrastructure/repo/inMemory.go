package repo

import "errors"

type InMemoryDb struct {
	Data map[string]EventDb
}

func NewInMemoryDb() InMemoryDb {
	data := make(map[string]EventDb)
	return InMemoryDb{Data: data}
}

func (db *InMemoryDb) GetOne(id string) (*EventDb, error) {
	if e, ok := db.Data[id]; ok {
		return &e, nil
	}
	return nil, nil
}

func (db *InMemoryDb) GetAll() []EventDb {
	events := make([]EventDb, len(db.Data))
	for _, v := range db.Data {
		events = append(events, v)
	}
	return events
}

func (db *InMemoryDb) Save(newEvent *EventDb) (*EventDb, error) {
	if _, ok := db.Data[newEvent.Id]; ok {
		return nil, errors.New("event already exists")
	}

	db.Data[newEvent.Id] = *newEvent
	return newEvent, nil
}
