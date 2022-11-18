package service

import (
	"github.com/ssharifzoda/levelup/internal/database"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

const (
	cholericID    = 41
	sangUniqID    = 40
	melancholicID = 43
	phlegmaticID  = 42
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
func (p *PublicService) DoTest(userId int, i domain.Test) (string, error) {
	ch := choleric(i) / znam(i)
	s := sangUniq(i) / znam(i)
	m := melancholic(i) / znam(i)
	ph := phlegmatic(i) / znam(i)
	var max int
	slice := []int{ch, s, m, ph}
	for i := range slice {
		max = slice[0]
		if slice[i] > max {
			max = slice[i]
		}
		continue
	}
	switch max {
	case ch:
		return p.db.DoTest(userId, cholericID)
	case s:
		return p.db.DoTest(userId, sangUniqID)
	case m:
		return p.db.DoTest(userId, melancholicID)
	case ph:
		return p.db.DoTest(userId, phlegmaticID)
	}
	return "", nil
}

func choleric(i domain.Test) int {
	res := (i.First + i.Second + i.Three + i.Fourth + i.Five) * 100
	return res
}
func sangUniq(i domain.Test) int {
	res := (i.Sixth + i.Seventh + i.Eighth + i.Ninth + i.Tenth) * 100
	return res
}
func phlegmatic(i domain.Test) int {
	res := (i.Eleventh + i.Twelfth + i.Thirteenth + i.Fourteenth + i.Fifteenth) * 100
	return res
}
func melancholic(i domain.Test) int {
	res := (i.Sixteenth + i.Seventeenth + i.Eighteenth + i.Nineteenth + i.Twentieth) * 100
	return res
}
func znam(i domain.Test) int {
	return choleric(i) + sangUniq(i) + phlegmatic(i) + melancholic(i)
}
