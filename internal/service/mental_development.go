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
	video    = "\\video.mp4"
	diet     = "\\diet.txt"
	plan     = "\\plan.txt"
	playlist = "playlist.rar"
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
	item.Audio = viper.GetString("storage.mentalcourse") + item.MentalCategory + audio
	if err != nil {
		return item, err
	}
	path, err := os.ReadFile(viper.GetString("storage.mentalcourse") + item.MentalCategory + filmPath)
	if err != nil {
		return item, err
	}
	item.FilmPath = string(path)
	item.Book = viper.GetString("storage.mentalcourse") + item.MentalCategory + book
	if err != nil {
		return item, err
	}
	return item, nil
}
func (m *MentalDevelopmentService) DeleteCourseById(userId, id int) (string, error) {
	return m.db.DeleteCourseById(userId, id)
}
func (m *MentalDevelopmentService) ValidateCategory(categoryId, userId int) (string, error) {
	return m.db.GetCategory(categoryId, userId)
}
func (m *MentalDevelopmentService) GetCategories() ([]domain.MentalCourseCategory, error) {
	return m.db.GetCategories()
}
