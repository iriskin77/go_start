package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Названия таблиц в константах
const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

// Здесь реализована логика подключения к БД, имена таблиц в const

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// Функция NewPostgresDB возвращает указатель на структуру sqlxDB
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	// С помощью функции Ping проверяется подключение
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
