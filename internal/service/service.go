package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	"github.com/ssharifzoda/levelup/internal/types"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	Validate(username, password string) (string, error)
}
type Diary interface {
	Create(userId int, item domain.Item) (int, error)
	GetAll(userId int) ([]domain.Item, error)
	GetById(userId, itemId int) (domain.Item, error)
	DeleteItemById(userId, itemId int) (string, error)
}
type BadHabit interface {
	Create(userId int, input domain.BadHabit) (int, error)
	GetAll(userId int) ([]domain.BadHabit, error)
	GetById(userId, id int) (domain.BadHabit, error)
	DeleteHabitById(userId, id int) (string, error)
	DoExercise(userId, id int, input domain.DoExercise) (string, error)
}

type Service struct {
	Authorization
	Diary
	BadHabit
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: NewAuthService(db.Authorization),
		Diary:         NewDiaryService(db.Diary),
		BadHabit:      NewBadHabitService(db.BadHabit),
	}
}
