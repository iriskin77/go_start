package repository

import (
	gostart "github.com/iriskin77/go_start"
	"github.com/jmoiron/sqlx"
)

// Интерфейсы называются в зависимости от участков доменной зоны, за которую они отвечают
type Authorization interface {
	CreateUser(user gostart.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// Конструктор сервисов
// Поскольку репозиторий должен работать с БД, то
func NewRepository(db *sqlx.DB) *Repository {
	// В файле репозитория инициализируем наш репозиторий в конструкторе
	return &Repository{Authorization: NewAuthPostgres(db)}
}
