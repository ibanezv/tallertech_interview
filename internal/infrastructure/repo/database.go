package repo

type DataBase interface {
	GetById(string) (*EventDb, error)
	GetAll() ([]EventDb, error)
	Save(*EventDb) (*EventDb, error)
}
