package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"time"
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
func (b *BadHabitService) GetAll(userId int) ([]domain.BadHabit, error) {
	return b.db.GetAll(userId)
}
func (b *BadHabitService) GetById(userId, id int) (domain.BadHabit, error) {
	return b.db.GetById(userId, id)
}
func (b *BadHabitService) DeleteHabitById(userId, id int) (string, error) {
	return b.db.DeleteHabitById(userId, id)
}
func (b *BadHabitService) DoExercise(userId, id int, input domain.DoExercise) (string, error) {
	input.LastSession = time.Now().String()
	input.Session = 1
	return b.db.DoExercise(userId, id, input)
}
