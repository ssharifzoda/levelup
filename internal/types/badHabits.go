package domain

import "time"

type BadHabitInput struct {
	Id              int       `json:"id"`
	HabitCategoryId int       `json:"habit_category_id" db:"habit_category_id" binding:"required"`
	EquivalentId    int       `json:"equivalent_id" db:"equivalent_id" binding:"required"`
	Created         time.Time `json:"created"`
}
type BadHabitsList struct {
	Id         int `json:"id" db:"id" binding:"required"`
	UserId     int `json:"user_id" db:"user_id" binding:"required"`
	BadHabitId int `json:"bad_habit_id" db:"bad_habit_id" binding:"required"`
}
type Exercise struct {
	BadHabitId int       `json:"bad_habit_id"`
	Registrar  int       `json:"registrar"`
	LastAt     time.Time `json:"last_at"`
}
type HabitsCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Equivalents struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type BadHabitOutput struct {
	Id            int       `json:"id"`
	HabitCategory string    `json:"habit_category" db:"name"`
	Equivalent    string    `json:"equivalent" db:"name"`
	Created       time.Time `json:"created"`
}
