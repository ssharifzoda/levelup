package domain

import "time"

type BadHabit struct {
	Id         int        `json:"id"`
	BadHabit   string     `json:"bad_habit" binding:"required"`
	Equivalent string     `json:"equivalent" binding:"required"`
	Status     string     `json:"status" binding:"required"`
	Session    time.Timer `json:"session"`
}
type BadHabitsList struct {
	Id         int `json:"id" db:"id" binding:"required"`
	UserId     int `json:"user_id" db:"user_id" binding:"required"`
	BadHabitId int `json:"bad_habit_id" db:"bad_habit_id" binding:"required"`
}
