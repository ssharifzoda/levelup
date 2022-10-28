package service

import (
	"github.com/spf13/viper"
	"github.com/ssharifzoda/levelup/internal/database"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

type PhysicianDevelopService struct {
	db database.PhysicianDevelopment
}

func NewPhysicianDevelopService(db database.PhysicianDevelopment) *PhysicianDevelopService {
	return &PhysicianDevelopService{db: db}
}

func (p *PhysicianDevelopService) Create(userId int, input domain.BodyCourseInput) (int, error) {
	return p.db.Create(userId, input)
}
func (p *PhysicianDevelopService) GetById(userId int, id int) (domain.BodyCourseOutput, error) {
	item, err := p.db.GetById(userId, id)
	if err != nil {
		return item, err
	}
	item.Video = viper.GetString("storage.bodycourse") + item.Level + "\\" + item.TrainCategory + video
	if err != nil {
		return item, err
	}
	item.TrainPlan = viper.GetString("storage.bodycourse") + item.Level + "\\" + item.TrainCategory + plan
	item.Diet = viper.GetString("storage.bodycourse") + item.Level + "\\" + item.TrainCategory + diet
	item.Playlist = viper.GetString("storage.bodycourse") + item.Level + "\\" + playlist
	return item, nil
}
