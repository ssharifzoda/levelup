package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	"github.com/ssharifzoda/levelup/internal/types"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateTokens(username, password string) (string, string, error)
	ParseToken(token string) (int, error)
	UserValidate(username, password string) (string, error)
	ParseRefreshToken(refreshToken string) (string, string, error)
}
type Diary interface {
	Create(userId int, item domain.Item) (int, error)
	GetAll(userId int) ([]domain.Item, error)
	GetById(userId, itemId int) (domain.Item, error)
	DeleteItemById(userId, itemId int) (string, error)
}
type BadHabit interface {
	Create(userId int, input domain.BadHabitInput) (int, error)
	GetAll(userId int) ([]domain.BadHabitOutput, error)
	GetById(userId, id int) (domain.BadHabitOutput, error)
	DeleteHabitById(userId, id int) (string, error)
	ValidateCategory(categoryId, userId int) (string, error)
}

type MentalDevelopment interface {
	Create(userId int, input domain.CourseInput) (int, error)
	GetById(userId int, id int) (domain.CourseOutput, error)
	DeleteCourseById(userId, id int) (string, error)
}

type Service struct {
	Authorization
	Diary
	BadHabit
	MentalDevelopment
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization:     NewAuthService(db.Authorization),
		Diary:             NewDiaryService(db.Diary),
		BadHabit:          NewBadHabitService(db.BadHabit),
		MentalDevelopment: NewMentalDevelopmentService(db.MentalDevelopment),
	}
}
