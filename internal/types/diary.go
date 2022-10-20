package domain

type ItemList struct {
	Id      int `json:"id" db:"id" binding:"required"`
	UserId  int `json:"user_id" db:"user_id" binding:"required"`
	DiaryId int `json:"diary_id" db:"diary_id" binding:"required"`
}

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Body        string `json:"body" binding:"required"`
}
