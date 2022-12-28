package models

type User struct {
	Name     string `json:"name" form:"name"`
	Id       int    `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Mobile   string `json:"mobile" form:"mobile"`
	Password string `json:"password" form:"password"`
}
