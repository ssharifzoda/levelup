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
	GetAll(userId, pageNo, itemLimit int) ([]domain.Item, error)
	GetById(userId, itemId int) (domain.Item, error)
	DeleteItemById(userId, itemId int) (string, error)
	GetItemByTitle(userId int, title string) (domain.Item, error)
}
type BadHabit interface {
	Create(userId int, input domain.BadHabitInput) (int, error)
	GetAll(userId int) ([]domain.BadHabitOutput, error)
	GetById(userId, id int) (domain.BadHabitOutput, error)
	DeleteHabitById(userId, id int) (string, error)
	ValidateCategory(categoryId, userId int) (string, error)
	GetCategories(pageNo, itemLimit int) ([]domain.HabitsCategory, error)
	GetEquivalents(pageNo, itemLimit int) ([]domain.Equivalents, error)
	EditEquivalentByID(userId, id, equivalent int) error
}

type MentalDevelopment interface {
	Create(userId int, input domain.CourseInput) (int, error)
	GetById(userId int, id int) (domain.CourseOutput, error)
	DeleteCourseById(userId, id int) (string, error)
	ValidateCategory(categoryId, userId int) (string, error)
	GetCategories() ([]domain.MentalCourseCategory, error)
}

type PhysicianDevelopment interface {
	Create(userId int, input domain.BodyCourseInput) (int, error)
	GetById(userId int, id int) (domain.BodyCourseOutput, error)
	DeleteCourseById(userId int, id int) (string, error)
	ValidateCategory(trainCategoryId, userId int) (string, error)
	GetCategories() ([]domain.BodyCourseCategories, error)
	GetLevels() ([]domain.BodyLevelCourse, error)
}

type Public interface {
	ReceivePublic(userId int, input domain.Public) error
	DoTest(userId int, input domain.Test) (string, error)
}

type Service struct {
	Authorization
	Diary
	BadHabit
	MentalDevelopment
	PhysicianDevelopment
	Public
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization:        NewAuthService(db.Authorization),
		Diary:                NewDiaryService(db.Diary),
		BadHabit:             NewBadHabitService(db.BadHabit),
		MentalDevelopment:    NewMentalDevelopmentService(db.MentalDevelopment),
		PhysicianDevelopment: NewPhysicianDevelopService(db.PhysicianDevelopment),
		Public:               NewPublicService(db),
	}
}
