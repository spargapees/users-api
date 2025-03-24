package dto

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type UserUpdate struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
