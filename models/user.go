package models

type User struct {
	Id       int
	Phone    string
	Email    string
	Password string
	Status   int
}

//表名 这个式内置的 gin框架回自动解析它  这个方法必须写
func (User) TableName() string {
	return "user"
}
