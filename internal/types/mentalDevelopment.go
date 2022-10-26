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
	Id             int    `json:"id" db:"id"`
	MentalCategory string `json:"mental_category" db:"mental_category_id"`
	//Audio          zip.File    `json:"audio"`
	FilmPath string `json:"film_path"`
	//Book           zip.File    `json:"book"`
	Created time.Time `json:"created" db:"created"`
}
