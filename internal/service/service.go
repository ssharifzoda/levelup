package service

import (
	"github.com/ssharifzoda/levelup/internal/domain"
	"github.com/ssharifzoda/levelup/pkg/database"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Diary interface {
}
type DiaryItems interface {
}

type Service struct {
	Authorization
	Diary
	DiaryItems
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: NewAuthService(db.Authorization),
	}
}
