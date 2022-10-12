package domain

type DiaryList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UserList struct {
	Id     int
	UserId int
	ListId int
}

type DiaryItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}
type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
