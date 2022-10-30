package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

type BadHabitService struct {
	db database.BadHabit
}

func NewBadHabitService(db database.BadHabit) *BadHabitService {
	return &BadHabitService{db: db}
}

func (b *BadHabitService) Create(userId int, input domain.BadHabitInput) (int, error) {
	return b.db.Create(userId, input)
}
func (b *BadHabitService) GetAll(userId int) ([]domain.BadHabitOutput, error) {
	return b.db.GetAll(userId)
}
func (b *BadHabitService) GetById(userId, id int) (domain.BadHabitOutput, error) {
	return b.db.GetById(userId, id)
}
func (b *BadHabitService) DeleteHabitById(userId, id int) (string, error) {
	return b.db.DeleteHabitById(userId, id)
}
func (b *BadHabitService) ValidateCategory(categoryId, userId int) (string, error) {
	return b.db.GetCategory(categoryId, userId)
}
func (b *BadHabitService) GetCategories() ([]domain.HabitsCategory, error) {
	return b.db.GetCategories()
}
func (b *BadHabitService) GetEquivalents() ([]domain.Equivalents, error) {
	return b.db.GetEquivalents()
}
