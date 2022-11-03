package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

type PublicService struct {
	db database.Public
}

func NewPublicService(db database.Public) *PublicService {
	return &PublicService{db: db}
}
func (p *PublicService) ReceivePublic(userId int, input domain.Public) error {
	return p.db.ReceivePublic(userId, input)
}
