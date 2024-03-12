package gostart

// Структура Юзера. Json теги нужны, чтобы корректно принимать данные от http запросов
type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
