package database

import (
	"github.com/jackc/pgx"
	"github.com/ssharifzoda/levelup/internal/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}
type Diary interface {
}
type DiaryItems interface {
}

type Database struct {
	Authorization
	Diary
	DiaryItems
}

func NewDatabase(conn *pgx.Conn) *Database {
	return &Database{
		Authorization: NewAuthPostgres(conn),
	}
}
