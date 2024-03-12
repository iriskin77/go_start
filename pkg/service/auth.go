package service

import (
	gostart "github.com/iriskin77/go_start"
	"github.com/iriskin77/go_start/pkg/repository"
)

type AuthService struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	repo repository.Authorization
}

func NewAuthService(repo repository.Repository) *AuthService {
	// Конструктор: принимает репозиторий, возваращает сервис с репозиторием
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user gostart.User) (int, error) {
	// Здесь передаем данные о пользователей еще на слой ниже, то есть в репозитрий,
	// где и будет сама работа с БД(запросы итд)
	return s.repo.CreateUser(user)
}
