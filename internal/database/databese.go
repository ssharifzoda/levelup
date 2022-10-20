package database

import (
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/types"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}
type Diary interface {
	Create(userId int, item domain.Item) (int, error)
	GetAll(userId int) ([]domain.Item, error)
	GetById(userId, itemId int) (domain.Item, error)
	DeleteItemById(userId, itemId int) (string, error)
}

type Database struct {
	Authorization
	Diary
}

func NewDatabase(conn *pgx.Conn) *Database {
	return &Database{
		Authorization: NewAuthPostgres(conn),
		Diary:         NewDiaryItemPostgres(conn),
	}
}
