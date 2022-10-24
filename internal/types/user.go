package domain

type User struct {
	Id       int    `json:"_" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}
