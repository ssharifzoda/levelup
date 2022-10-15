package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	"github.com/ssharifzoda/levelup/internal/domain"
)

type DiaryService struct {
	db database.Diary
}

func NewDiaryService(db database.Diary) *DiaryService {
	return &DiaryService{db: db}
}
func (d *DiaryService) Create(userId int, item domain.Item) (int, error) {
	return d.db.Create(userId, item)
}
