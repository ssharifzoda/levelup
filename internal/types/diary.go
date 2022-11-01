package domain

type UsersSpace struct {
	UserId  int `json:"user_id" db:"user_id" binding:"required"`
	DiaryId int `json:"item_id" db:"item_id" binding:"required"`
}

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Body        string `json:"body" binding:"required"`
	Created     string `json:"created"`
}
