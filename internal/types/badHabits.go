package domain

type BadHabit struct {
	Id         int    `json:"id"`
	BadHabit   string `json:"bad_habit" binding:"required"`
	Equivalent string `json:"equivalent" binding:"required"`
	Session    int    `json:"session"`
}
type BadHabitsList struct {
	Id         int `json:"id" db:"id" binding:"required"`
	UserId     int `json:"user_id" db:"user_id" binding:"required"`
	BadHabitId int `json:"bad_habit_id" db:"bad_habit_id" binding:"required"`
}
type DoExercise struct {
	Session     int    `json:"session" binding:"required"`
	LastSession string `json:"last_session"`
}
