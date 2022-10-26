package service

import (
	"github.com/spf13/viper"
	"github.com/ssharifzoda/levelup/internal/database"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"os"
)

const (
	audio    = "audio.ogg"
	slesh    = "\\"
	filmPath = "filmpath.txt"
	book     = "book.epub"
)

type MentalDevelopmentService struct {
	db database.MentalDevelopment
}

func NewMentalDevelopmentService(db database.MentalDevelopment) *MentalDevelopmentService {
	return &MentalDevelopmentService{db: db}
}

func (m *MentalDevelopmentService) Create(userId int, input domain.CourseInput) (int, error) {
	return m.db.Create(userId, input)
}
func (m *MentalDevelopmentService) GetById(userId int, id int) (domain.CourseOutput, error) {
	item, err := m.db.GetById(userId, id)
	if err != nil {
		return item, err
	}
	item.Audio = viper.GetString("storage.confidence") + audio
	if err != nil {
		return item, err
	}
	path, err := os.ReadFile(viper.GetString("storage.confidence") + filmPath)
	if err != nil {
		return item, err
	}
	item.FilmPath = string(path)
	item.Book = viper.GetString("storage.confidence") + book
	if err != nil {
		return item, err
	}
	return item, nil
}
