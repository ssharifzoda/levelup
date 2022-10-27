package service

import (
	"github.com/spf13/viper"
	"github.com/ssharifzoda/levelup/internal/database"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"os"
)

const (
	audio    = "\\audio.ogg"
	filmPath = "\\filmpath.txt"
	book     = "\\book.epub"
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
	item.Audio = viper.GetString("storage") + item.MentalCategory + audio
	if err != nil {
		return item, err
	}
	path, err := os.ReadFile(viper.GetString("storage") + item.MentalCategory + filmPath)
	if err != nil {
		return item, err
	}
	item.FilmPath = string(path)
	item.Book = viper.GetString("storage") + item.MentalCategory + book
	if err != nil {
		return item, err
	}
	return item, nil
}
func (m *MentalDevelopmentService) DeleteCourseById(userId, id int) (string, error) {
	return m.db.DeleteCourseById(userId, id)
}
