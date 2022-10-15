package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	"github.com/ssharifzoda/levelup/internal/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Diary interface {
	Create(userId int, item domain.Item) (int, error)
}

type Service struct {
	Authorization
	Diary
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: NewAuthService(db.Authorization),
		Diary:         NewDiaryService(db.Diary),
	}
}
