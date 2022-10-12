package service

import "github.com/ssharifzoda/levelup/pkg/database"

type Authorization interface {
}
type Diary interface {
}
type DiaryItems interface {
}

type Service struct {
	Authorization
	Diary
	DiaryItems
}

func NewService(db *database.Database) *Service {
	return &Service{}
}