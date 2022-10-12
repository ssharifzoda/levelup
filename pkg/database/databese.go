package database

type Authorization interface {
}
type Diary interface {
}
type DiaryItems interface {
}

type Database struct {
	Authorization
	Diary
	DiaryItems
}

func NewDatabase() *Database {
	return &Database{}
}
