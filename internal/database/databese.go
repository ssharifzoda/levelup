package database

import (
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/types"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
	GetUserById(userId int) (string, string, error)
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
	GetCategory(categoryId, userId int) (string, error)
}

type MentalDevelopment interface {
	Create(userId int, input domain.CourseInput) (int, error)
	GetById(userId int, id int) (domain.CourseOutput, error)
	DeleteCourseById(userId, id int) (string, error)
	GetCategory(categoryId, userId int) (string, error)
}

type PhysicianDevelopment interface {
	Create(userId int, input domain.BodyCourseInput) (int, error)
	GetById(userId int, id int) (domain.BodyCourseOutput, error)
	DeleteCourseById(userId, id int) (string, error)
	GetCategory(trainCategoryId, userId int) (string, error)
}

type Database struct {
	Authorization
	Diary
	BadHabit
	MentalDevelopment
	PhysicianDevelopment
}

func NewDatabase(conn *pgx.Conn) *Database {
	return &Database{
		Authorization:        NewAuthPostgres(conn),
		Diary:                NewDiaryItemPostgres(conn),
		BadHabit:             NewBadHabitPostgres(conn),
		MentalDevelopment:    NewMentalDevelopPostgres(conn),
		PhysicianDevelopment: NewPhysicianDevelopPostgres(conn),
	}
}
