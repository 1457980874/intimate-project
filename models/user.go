package models

type UserModels struct {
	Name string `form:"username"`
	Phone string `form:"phone"`
	PassWord string `form:"password"`
	IDcard string `form:IDcard`
}


