package service

import (
	gostart "github.com/iriskin77/go_start"
	"github.com/iriskin77/go_start/pkg/repository"
)

// Интерфейсы называются в зависимости от участков доменной зоны, за которую они отвечают

type Authorization interface {
	CreateUser(user gostart.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// Конструктор сервисов. Сервисы будут передавать данные из хэндлера ниже, на уровень репозитория, поэтому нужен указатель
// на структуру репозитория (репозиторий коннектиться к БД)

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
