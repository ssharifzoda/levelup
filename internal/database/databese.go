package database

import (
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/types"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
	GetUserById(userId int) (string, string, error)
}
type Diary interface {
	Create(userId int, item domain.Item) (int, error)
	GetAll(userId, offset, itemLimit int) ([]domain.Item, error)
	GetById(userId, itemId int) (domain.Item, error)
	DeleteItemById(userId, itemId int) (string, error)
	GetItemByTitle(userId int, title string) (domain.Item, error)
}

type BadHabit interface {
	Create(userId int, input domain.BadHabitInput) (int, error)
	GetAll(userId int) ([]domain.BadHabitOutput, error)
	GetById(userId, id int) (domain.BadHabitOutput, error)
	DeleteHabitById(userId, id int) (string, error)
	GetCategory(categoryId, userId int) (string, error)
	GetCategories(offset, itemLimit int) ([]domain.HabitsCategory, error)
	GetEquivalents(offset, itemLimit int) ([]domain.Equivalents, error)
	EditEquivalentByID(userId, id, equivalent int) error
}

type MentalDevelopment interface {
	Create(userId int, input domain.CourseInput) (int, error)
	GetById(userId int, id int) (domain.CourseOutput, error)
	DeleteCourseById(userId, id int) (string, error)
	GetCategory(categoryId, userId int) (string, error)
	GetCategories() ([]domain.MentalCourseCategory, error)
}

type PhysicianDevelopment interface {
	Create(userId int, input domain.BodyCourseInput) (int, error)
	GetById(userId int, id int) (domain.BodyCourseOutput, error)
	DeleteCourseById(userId, id int) (string, error)
	GetCategory(trainCategoryId, userId int) (string, error)
	GetCategories() ([]domain.BodyCourseCategories, error)
	GetLevels() ([]domain.BodyLevelCourse, error)
}

type Public interface {
	ReceivePublic(userId int, input domain.Public) error
}

type Database struct {
	Authorization
	Diary
	BadHabit
	MentalDevelopment
	PhysicianDevelopment
	Public
}

func NewDatabase(conn *pgx.Conn, session *gorm.DB) *Database {
	return &Database{
		Authorization:        NewAuthPostgres(conn),
		Diary:                NewDiaryItemPostgres(conn, session),
		BadHabit:             NewBadHabitPostgres(conn, session),
		MentalDevelopment:    NewMentalDevelopPostgres(conn),
		PhysicianDevelopment: NewPhysicianDevelopPostgres(conn),
		Public:               NewPublicPostgres(conn, session),
	}
}
