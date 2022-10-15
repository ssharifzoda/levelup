package domain

type ItemList struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	DiaryId int `json:"diary_id"`
}

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Body        string `json:"body" binding:"required"`
}
