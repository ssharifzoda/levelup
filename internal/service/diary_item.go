package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	"github.com/ssharifzoda/levelup/internal/types"
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
func (d *DiaryService) GetAll(userId int) ([]domain.Item, error) {
	return d.db.GetAll(userId)
}
func (d *DiaryService) GetById(userId, itemId int) (domain.Item, error) {
	return d.db.GetById(userId, itemId)

}
func (d *DiaryService) DeleteItemById(userId, itemId int) (string, error) {
	return d.db.DeleteItemById(userId, itemId)
}
func (d *DiaryService) GetItemByTitle(userId int, title string) (domain.Item, error) {
	return d.db.GetItemByTitle(userId, title)
}
