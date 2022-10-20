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

func (b *BadHabitService) Create(userId int, input domain.BadHabit) (int, error) {
	return b.db.Create(userId, input)
}
