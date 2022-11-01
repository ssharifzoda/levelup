package domain

import "time"

type BodyCourseInput struct {
	Id              int       `json:"id" db:"id"`
	LevelId         int       `json:"level_id" db:"level_category_id" binding:"required"`
	TrainCategoryId int       `json:"train_category_id" db:"train_category_id" binding:"required"`
	Created         time.Time `json:"created" db:"created"`
}
type BodyCourseOutput struct {
	Id            int       `json:"id" db:"id"`
	Level         string    `json:"Level" db:"name"`
	TrainCategory string    `json:"Train-category" db:"name"`
	TrainPlan     string    `json:"Train-plan"`
	Video         string    `json:"Video"`
	Playlist      string    `json:"Playlist"`
	Diet          string    `json:"Diet"`
	Created       time.Time `json:"created" db:"created"`
}
type BodyCourseCategories struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type BodyLevelCourse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
