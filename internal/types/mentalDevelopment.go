package domain

import (
	"time"
)

type CourseInput struct {
	Id               int       `json:"id" db:"id"`
	MentalCategoryId int       `json:"mental_category_id" db:"mental_category_id" binding:"required"`
	Created          time.Time `json:"created" db:"created"`
}
type CourseOutput struct {
	Id             int       `json:"id" db:"id"`
	MentalCategory string    `json:"category" db:"mental_category_id"`
	Audio          string    `json:"audio-book"`
	FilmPath       string    `json:"film-path"`
	Book           string    `json:"book-path"`
	Created        time.Time `json:"created" db:"created"`
}
